// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package pointerdb

import (
	//"bytes"
	"context"
	"fmt"
	"errors"
	"testing"
	"log"
	"strings"

	"github.com/golang/protobuf/proto"
	//"github.com/spf13/viper"
	//"go.uber.org/zap"
	//"google.golang.org/grpc"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	pb "storj.io/storj/protos/pointerdb"
	p "storj.io/storj/pkg/paths"
)

const (
	unauthenticated = "failed API creds"
	noPathGiven = "file path not given"
)

var (
	ctx = context.Background()
	ErrUnauthenticated = errors.New(unauthenticated)
	ErrNoFileGiven = errors.New(noPathGiven)
)

func TestNewNetStateClient(t *testing.T) {
	// mocked grpcClient so we don't have
	// to call the network to test the code
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	gc:= NewMockNetStateClient(ctrl)
	nsc := NetState{grpcClient: gc}

	assert.NotNil(t, nsc)
	assert.NotNil(t, nsc.grpcClient)
}


func MakePointer(path p.Path, auth []byte) pb.PutRequest {
	// rps is an example slice of RemotePieces to add to this
	// REMOTE pointer type.
	var rps []*pb.RemotePiece
	rps = append(rps, &pb.RemotePiece{
		PieceNum: int64(1),
		NodeId:   "testId",
	})
	pr := pb.PutRequest{
		Path: path.Bytes(),
		Pointer: &pb.Pointer{
			Type: pb.Pointer_REMOTE,
			Remote: &pb.RemoteSegment{
				Redundancy: &pb.RedundancyScheme{
					Type:             pb.RedundancyScheme_RS,
					MinReq:           int64(1),
					Total:            int64(3),
					RepairThreshold:  int64(2),
					SuccessThreshold: int64(3),
				},
				PieceId:      "testId",
				RemotePieces: rps,
			},
			Size: int64(1),
		},
		APIKey: auth,
	}
	return pr
}

func TestPut(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for i, tt := range []struct {
		APIKey []byte
		path p.Path
		err error 
		errString string
	}{
		{[]byte("abc123"), p.New("file1/file2"), nil, ""},
		{[]byte("wrong key"), p.New("file1/file2"), ErrUnauthenticated,unauthenticated},
		{[]byte("abc123"), p.New(""), ErrNoFileGiven, noPathGiven},
		{[]byte("wrong key"), p.New(""), ErrUnauthenticated, unauthenticated},
		{[]byte(""), p.New(""), ErrUnauthenticated, unauthenticated},
	}{
		pr:= MakePointer(tt.path, tt.APIKey)

		errTag := fmt.Sprintf("Test case #%d", i)
		gc:= NewMockNetStateClient(ctrl)
		nsc := NetState{grpcClient: gc}

		gomock.InOrder(
			gc.EXPECT().Put(ctx, &pr).Return(nil, tt.err),
		)

		err := nsc.Put(ctx, tt.path, pr.Pointer, tt.APIKey)
		
		if err != nil {
			assert.EqualError(t, err, tt.errString, errTag)
		} else {
			assert.NoError(t, err, errTag)
		}
	}
}

func TestGet(t *testing.T){
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for i, tt := range []struct {
		APIKey []byte
		path p.Path
		err error 
		errString string
	}{
		{[]byte("wrong key"), p.New("file1/file2"), ErrUnauthenticated,unauthenticated},
		{[]byte("abc123"), p.New(""), ErrNoFileGiven, noPathGiven},
		{[]byte("wrong key"), p.New(""), ErrUnauthenticated, unauthenticated},
		{[]byte(""), p.New(""), ErrUnauthenticated, unauthenticated},
		{[]byte("abc123"), p.New("file1/file2"), nil, ""},
	}{
		pr:= MakePointer(tt.path, tt.APIKey)
		gr:= pb.GetRequest{Path: tt.path.Bytes(), APIKey: tt.APIKey}
		
		data, err := proto.Marshal(pr.Pointer)
		if err != nil {
			log.Fatal("marshaling error: ", err)
		}

		byteData := []byte(data)

		grr := pb.GetResponse{byteData}

		errTag := fmt.Sprintf("Test case #%d", i)
		
		gc:= NewMockNetStateClient(ctrl)
		nsc := NetState{grpcClient: gc}

		gomock.InOrder(
			gc.EXPECT().Get(ctx, &gr).Return(&grr, tt.err),
		)

		pointer, err := nsc.Get(ctx, tt.path, tt.APIKey)

		fmt.Println("pointer is: ", pointer)

		if err != nil {
			assert.EqualError(t, err, tt.errString, errTag)
		} else {
			assert.NoError(t, err, errTag)
		}
	}
}

func TestList(t *testing.T){

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	for i, tt := range []struct {
		APIKey []byte
		startingPath p.Path
		limit int64 
		truncated bool
		paths []string
		err error 
		errString string
	}{
		{[]byte("wrong key"), p.New(""), 2, true, []string{""}, ErrUnauthenticated, unauthenticated},
		{[]byte("abc123"), p.New("file1/file2"), 2, true, []string{"test"},  nil, ""},
		{[]byte("abc123"), p.New(""), 2, true, []string{"file1/file2", "file3/file4", "file1", "file1/file2/great3", "test"},  ErrNoFileGiven, noPathGiven},
		{[]byte("abc123"), p.New("file1/file2"), 2, false, []string{""},  nil, ""},
		{[]byte("wrong key"), p.New("file1/file2"), 2, true, []string{"file1/file2", "file3/file4", "file1", "file1/file2/great3", "test"}, ErrUnauthenticated,unauthenticated},
		{[]byte("abc123"), p.New("file1/file2"), 3, true, []string{"file1/file2", "file3/file4", "file1", "file1/file2/great3", "test"},  nil, ""},
	}{
		lr := pb.ListRequest{
			StartingPathKey: tt.startingPath.Bytes(),
			Limit:           tt.limit,
			APIKey:          tt.APIKey,
		}

		var truncatedPathsBytes [][]byte

		getCorrectPaths := func(fileName string) bool { return strings.HasPrefix(fileName, "file1")}
		filterPaths := filterPathName(tt.paths, getCorrectPaths)
		
		if len(filterPaths) == 0 {
			truncatedPathsBytes = [][]byte{{}}
		} else{
			truncatedPaths := filterPaths[0:tt.limit]
			truncatedPathsBytes := make([][]byte, len(truncatedPaths))
		
			for i, pathName := range truncatedPaths {
				bytePathName := []byte(pathName)
				truncatedPathsBytes[i] = make([]byte, 1)
				truncatedPathsBytes[i] = bytePathName 
			}
		}
			
		lrr := pb.ListResponse{Paths: truncatedPathsBytes, Truncated: tt.truncated }

		errTag := fmt.Sprintf("Test case #%d", i)

		gc:= NewMockNetStateClient(ctrl)
		nsc := NetState{grpcClient: gc}

		gomock.InOrder(
			gc.EXPECT().List(ctx, &lr).Return(&lrr, tt.err),
		)

		paths, trunc, err := nsc.List(ctx, tt.startingPath,  tt.limit, tt.APIKey)
		
		if err != nil {
			assert.EqualError(t, err, tt.errString, errTag)
		} else {
			assert.NoError(t, err, errTag)
		}

		fmt.Println("Path is: ", paths, "Trunc is: ",  trunc, "Err is: ", err)
	}
}

func filterPathName(pathString []string, test func(string) bool) (filteredPathNames []string) {
	for _, name := range pathString{
		if test(name) {
			filteredPathNames = append(filteredPathNames, name)
		}
	}
	return
}


func TestDelete(t *testing.T){

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

}



// func TestList(t *testing.T) {
// 	// nt := NewNetStateClientTest(t)
// 	// defer nt.Close()
// 	mockNetStateService := new(MockedNetState)

// 	reqs := MakePointers(4)
// 	for _, req := range reqs {
// 		mockNetStateService.Put(req)
// 	}

// 	listReq := pb.ListRequest{
// 		StartingPathKey: []byte("file/path/2"),
// 		Limit:           5,
// 		APIKey:          []byte("abc123"),
// 	}
// 	listRes := nt.List(listReq)
// 	if listRes.Truncated {
// 		nt.HandleErr(nil, "Expected list slice to not be truncated")
// 	}
// 	if !bytes.Equal(listRes.Paths[0], []byte("file/path/2")) {
// 		nt.HandleErr(nil, "Failed to list correct file paths")
// 	}
// }







// func (nt *NetStateClientTest) Put(pr pb.PutRequest) *pb.PutResponse {
// 	pre := nt.mdb.PutCalled
// 	putRes, err := nt.c.Put(ctx, &pr)
// 	if err != nil {
// 		nt.HandleErr(err, "Failed to put")
// 	}
// 	if pre+1 != nt.mdb.PutCalled {
// 		nt.HandleErr(nil, "Failed to call Put correct number of times")
// 	}
// 	return putRes
// }


// func (m *MockedNetState) Get(ctx context.Context, path p.Path, APIKey []byte) (*pb.Pointer, error) {
// 	args := m.Called(ctx, path, APIKey)
// 	pre := m.mdb.GetCalled

// 	_, err := m.c.Get(ctx, &pb.GetRequest{Path: path.Bytes(), APIKey: APIKey})
	
// 	if err != nil {
// 		m.HandleErr(err, "Failed to get")
// 	}

// 	if pre+1 != m.mdb.GetCalled {
// 		m.HandleErr(nil, "Failed to call Get correct number of times")
// 	}
// 	return  args.Get(0).(*pb.Pointer), args.Error(1)
// }



// func (nt *NetStateClientTest) Get(gr pb.GetRequest) *pb.GetResponse {
// 	pre := nt.mdb.GetCalled
// 	getRes, err := nt.c.Get(ctx, &gr)
// 	if err != nil {
// 		nt.HandleErr(err, "Failed to get")
// 	}
// 	if pre+1 != nt.mdb.GetCalled {
// 		nt.HandleErr(nil, "Failed to call Get correct number of times")
// 	}
// 	return getRes
// }

// func (nt *NetStateClientTest) List(lr pb.ListRequest) (listRes *pb.ListResponse) {
// 	pre := nt.mdb.ListCalled
// 	listRes, err := nt.c.List(ctx, &lr)
// 	if err != nil {
// 		nt.HandleErr(err, "Failed to list")
// 	}
// 	if pre+1 != nt.mdb.ListCalled {
// 		nt.HandleErr(nil, "Failed to call List correct number of times")
// 	}
// 	return listRes
// }

// func (nt *NetStateClientTest) Delete(dr pb.DeleteRequest) (delRes *pb.DeleteResponse) {
// 	pre := nt.mdb.DeleteCalled
// 	delRes, err := nt.c.Delete(ctx, &dr)
// 	if err != nil {
// 		nt.HandleErr(err, "Failed to delete")
// 	}
// 	if pre+1 != nt.mdb.DeleteCalled {
// 		nt.HandleErr(nil, "Failed to call Delete correct number of times")
// 	}

// 	return delRes
// }


// func (m *MockedNetState) HandleErr(err error, msg string) {
// 	if err != nil {
// 		panic(err)
// 	}
// 	panic(msg)
// }


// func (nt *NetStateClientTest) HandleErr(err error, msg string) {
// 	nt.Error(msg)
// 	if err != nil {
// 		panic(err)
// 	}
// 	panic(msg)
// }

// func TestMockList(t *testing.T) {
// 	nt := NewNetStateClientTest(t)
// 	defer nt.Close()

// 	err := nt.mdb.Put([]byte("k1"), []byte("v1"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = nt.mdb.Put([]byte("k2"), []byte("v2"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = nt.mdb.Put([]byte("k3"), []byte("v3"))
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = nt.mdb.Put([]byte("k4"), []byte("v4"))
// 	if err != nil {
// 		panic(err)
// 	}

// 	keys, err := nt.mdb.List([]byte("k2"), 2)
// 	if err != nil {
// 		nt.HandleErr(err, "Failed to list")
// 	}
// 	if fmt.Sprintf("%s", keys) != "[k2 k3]" {
// 		nt.HandleErr(nil, "Failed to receive accepted list. Received "+fmt.Sprintf("%s", keys))
// 	}

// 	keys, err = nt.mdb.List(nil, 3)
// 	if err != nil {
// 		nt.HandleErr(err, "Failed to list")
// 	}
// 	if fmt.Sprintf("%s", keys) != "[k1 k2 k3]" {
// 		nt.HandleErr(nil, "Failed to receive accepted list. Received "+fmt.Sprintf("%s", keys))
// 	}
// }




// func TestNetStatePutGet(t *testing.T) {
// 	mockNetStateService := new(MockedNetState)

	//path := p.New("fold1/fold2/fold3/file.txt")
	//pr := MakePointer(path, true)
	//mockNetStateService.Put(ctx, path, pr.Pointer, pr.APIKey)

	// mockNetStateService.On("Put", ctx, path, pr.Pointer, []byte("abc123")).Return(nil)
	// mockNetStateService.On("Get", ctx, path, []byte("abc123")).Return(pr.Pointer, nil)

	// assert.Equal(t, pr.Pointer, pr.Pointer, "they should be equal")
	
	
	//(t, process.Main(func() error { return nil }, mockService))
//	mockNetStateService.AssertExpectations(t)

	// preGet := mockNetStateService.mdb.GetCalled
	// prePut := mockNetStateService.mdb.PutCalled

	//get fails here 
	//pointerA, err := mockNetStateService.Get(ctx, p.New("file/path/1"), []byte("abc123"))
	
	// if pointerA != nil {
	// 	mockNetStateService.HandleErr(nil, "Expected no pointer")
	// }
	
	// path := p.New("fold1/fold2/fold3/file.txt")
	
	// pr := MakePointer(path, true)
	// mockNetStateService.Put(ctx, path, pr.Pointer, pr.APIKey)

	
	// pointerB, err := mockNetStateService.Get(ctx, path,[]byte("abc123"))
	// if err != nil {
	// 	mockNetStateService.HandleErr(nil, "Failed to get the put pointer")
	// }

	// pointerBytes, err := proto.Marshal(pr.Pointer)
	
	// if err != nil {
	// 	mockNetStateService.HandleErr(err, "Failed to marshal test pointer")
	// }

	// if !bytes.Equal(pointerB, pointerBytes) {
	// 	mockNetStateService.HandleErr(nil, "Expected to get same content that was put")
	// }

	// if mockNetStateService.mdb.GetCalled != preGet+2 {
	// 	mockNetStateService.HandleErr(nil, "Failed to call get correct number of times")
	// }

	// if mockNetStateService.mdb.PutCalled != prePut+1 {
	// 	mockNetStateService.HandleErr(nil, "Failed to call put correct number of times")
	// }
//}





// func TestGetAuth(t *testing.T) {
// 	mockNetStateService := new(MockedNetState)

// 	_, err := mockNetStateService.Get(ctx, p.New("file/path/1"), []byte("wrong key"))
// 	mockNetStateService.On("Get", ctx, p.New("file/path/1"),[]byte("wrong key")).Return(nil, err)

// 	if err == nil {
// 		mockNetStateService.HandleErr(nil, "Failed to Get because of wrong auth key")
// 	}

// 	mockNetStateService.AssertExpectations(t)
// }

// func TestPutAuth(t *testing.T) {
// 	mockNetStateService := new(MockedNetState)

// 	path := p.New("file/path")
// 	pr := MakePointer(path, false)

// 	err := mockNetStateService.Put(ctx, path, pr.Pointer, pr.APIKey)
// 	mockNetStateService.On("Put", ctx, path, pr.Pointer, pr.APIKey).Return(nil, err)

// 	if err == nil {
// 		mockNetStateService.HandleErr(nil, "Failed to error for wrong auth key")
// 	}
// 	mockNetStateService.AssertExpectations(t)
// }

// func TestDelete(t *testing.T) {
// 	nt := NewNetStateClientTest(t)
// 	defer nt.Close()

// 	pre := nt.mdb.DeleteCalled

// 	reqs := MakePointers(1)
// 	_, err := nt.c.Put(ctx, &reqs[0])
// 	if err != nil {
// 		nt.HandleErr(err, "Failed to put")
// 	}

// 	delReq := pb.DeleteRequest{
// 		Path:   []byte("file/path/1"),
// 		APIKey: []byte("abc123"),
// 	}
// 	_, err = nt.c.Delete(ctx, &delReq)
// 	if err != nil {
// 		nt.HandleErr(err, "Failed to delete")
// 	}
// 	if pre+1 != nt.mdb.DeleteCalled {
// 		nt.HandleErr(nil, "Failed to call Delete correct number of times")
// 	}
// }

// func TestDeleteAuth(t *testing.T) {
// 	nt := NewNetStateClientTest(t)
// 	defer nt.Close()

// 	reqs := MakePointers(1)
// 	_, err := nt.c.Put(ctx, &reqs[0])
// 	if err != nil {
// 		nt.HandleErr(err, "Failed to put")
// 	}

// 	delReq := pb.DeleteRequest{
// 		Path:   []byte("file/path/1"),
// 		APIKey: []byte("wrong key"),
// 	}
// 	_, err = nt.c.Delete(ctx, &delReq)
// 	if err == nil {
// 		nt.HandleErr(nil, "Failed to error with wrong auth key")
// 	}
// }

// func TestList(t *testing.T) {
// 	// nt := NewNetStateClientTest(t)
// 	// defer nt.Close()
// 	mockNetStateService := new(MockedNetState)

// 	reqs := MakePointers(4)
// 	for _, req := range reqs {
// 		mockNetStateService.Put(req)
// 	}

// 	listReq := pb.ListRequest{
// 		StartingPathKey: []byte("file/path/2"),
// 		Limit:           5,
// 		APIKey:          []byte("abc123"),
// 	}
// 	listRes := nt.List(listReq)
// 	if listRes.Truncated {
// 		nt.HandleErr(nil, "Expected list slice to not be truncated")
// 	}
// 	if !bytes.Equal(listRes.Paths[0], []byte("file/path/2")) {
// 		nt.HandleErr(nil, "Failed to list correct file paths")
// 	}
// }

// func TestListTruncated(t *testing.T) {
// 	nt := NewNetStateClientTest(t)
// 	defer nt.Close()

// 	reqs := MakePointers(3)
// 	for _, req := range reqs {
// 		_, err := nt.c.Put(ctx, &req)
// 		if err != nil {
// 			nt.HandleErr(err, "Failed to put")
// 		}
// 	}

// 	listReq := pb.ListRequest{
// 		StartingPathKey: []byte("file/path/1"),
// 		Limit:           1,
// 		APIKey:          []byte("abc123"),
// 	}
// 	listRes, err := nt.c.List(ctx, &listReq)
// 	if err != nil {
// 		nt.HandleErr(err, "Failed to list file paths")
// 	}
// 	if !listRes.Truncated {
// 		nt.HandleErr(nil, "Expected list slice to be truncated")
// 	}
// }

// func TestListWithoutStartingKey(t *testing.T) {
// 	nt := NewNetStateClientTest(t)
// 	defer nt.Close()

// 	reqs := MakePointers(3)
// 	for _, req := range reqs {
// 		_, err := nt.c.Put(ctx, &req)
// 		if err != nil {
// 			nt.HandleErr(err, "Failed to put")
// 		}
// 	}

// 	listReq := pb.ListRequest{
// 		Limit:  3,
// 		APIKey: []byte("abc123"),
// 	}
// 	listRes, err := nt.c.List(ctx, &listReq)
// 	if err != nil {
// 		nt.HandleErr(err, "Failed to list without starting key")
// 	}

// 	if !bytes.Equal(listRes.Paths[2], []byte("file/path/3")) {
// 		nt.HandleErr(nil, "Failed to list correct paths")
// 	}
// }

// func TestListWithoutLimit(t *testing.T) {
// 	nt := NewNetStateClientTest(t)
// 	defer nt.Close()

// 	listReq := pb.ListRequest{
// 		StartingPathKey: []byte("file/path/3"),
// 		APIKey:          []byte("abc123"),
// 	}
// 	_, err := nt.c.List(ctx, &listReq)
// 	if err == nil {
// 		t.Error("Failed to error when not given limit")
// 	}
// }

// func TestListAuth(t *testing.T) {
// 	nt := NewNetStateClientTest(t)
// 	defer nt.Close()

// 	listReq := pb.ListRequest{
// 		StartingPathKey: []byte("file/path/3"),
// 		Limit:           1,
// 		APIKey:          []byte("wrong key"),
// 	}
// 	_, err := nt.c.List(ctx, &listReq)
// 	if err == nil {
// 		t.Error("Failed to error when given wrong auth key")
// 	}
// }
