CREATE TABLE accounting_rollups (
	id bigserial NOT NULL,
	node_id bytea NOT NULL,
	start_time timestamp with time zone NOT NULL,
	put_total bigint NOT NULL,
	get_total bigint NOT NULL,
	get_audit_total bigint NOT NULL,
	get_repair_total bigint NOT NULL,
	put_repair_total bigint NOT NULL,
	at_rest_total double precision NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE accounting_timestamps (
	name text NOT NULL,
	value timestamp with time zone NOT NULL,
	PRIMARY KEY ( name )
);
CREATE TABLE bucket_bandwidth_rollups (
	bucket_name bytea NOT NULL,
	project_id bytea NOT NULL,
	interval_start timestamp NOT NULL,
	interval_seconds integer NOT NULL,
	action integer NOT NULL,
	inline bigint NOT NULL,
	allocated bigint NOT NULL,
	settled bigint NOT NULL,
	PRIMARY KEY ( bucket_name, project_id, interval_start, action )
);
CREATE TABLE bucket_storage_tallies (
	bucket_name bytea NOT NULL,
	project_id bytea NOT NULL,
	interval_start timestamp NOT NULL,
	inline bigint NOT NULL,
	remote bigint NOT NULL,
	remote_segments_count integer NOT NULL,
	inline_segments_count integer NOT NULL,
	object_count integer NOT NULL,
	metadata_size bigint NOT NULL,
	PRIMARY KEY ( bucket_name, project_id, interval_start )
);
CREATE TABLE bucket_usages (
	id bytea NOT NULL,
	bucket_id bytea NOT NULL,
	rollup_end_time timestamp with time zone NOT NULL,
	remote_stored_data bigint NOT NULL,
	inline_stored_data bigint NOT NULL,
	remote_segments integer NOT NULL,
	inline_segments integer NOT NULL,
	objects integer NOT NULL,
	metadata_size bigint NOT NULL,
	repair_egress bigint NOT NULL,
	get_egress bigint NOT NULL,
	audit_egress bigint NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE certRecords (
	publickey bytea NOT NULL,
	id bytea NOT NULL,
	update_at timestamp with time zone NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE injuredsegments (
	path text NOT NULL,
	data bytea NOT NULL,
	attempted timestamp,
	PRIMARY KEY ( path )
);
CREATE TABLE irreparabledbs (
	segmentpath bytea NOT NULL,
	segmentdetail bytea NOT NULL,
	pieces_lost_count bigint NOT NULL,
	seg_damaged_unix_sec bigint NOT NULL,
	repair_attempt_count bigint NOT NULL,
	PRIMARY KEY ( segmentpath )
);
CREATE TABLE nodes (
	id bytea NOT NULL,
	address text NOT NULL,
	last_ip text NOT NULL,
	protocol integer NOT NULL,
	type integer NOT NULL,
	email text NOT NULL,
	wallet text NOT NULL,
	free_bandwidth bigint NOT NULL,
	free_disk bigint NOT NULL,
	major bigint NOT NULL,
	minor bigint NOT NULL,
	patch bigint NOT NULL,
	hash text NOT NULL,
	timestamp timestamp with time zone NOT NULL,
	release boolean NOT NULL,
	latency_90 bigint NOT NULL,
	audit_success_count bigint NOT NULL,
	total_audit_count bigint NOT NULL,
	audit_success_ratio double precision NOT NULL,
	uptime_success_count bigint NOT NULL,
	total_uptime_count bigint NOT NULL,
	uptime_ratio double precision NOT NULL,
	created_at timestamp with time zone NOT NULL,
	updated_at timestamp with time zone NOT NULL,
	last_contact_success timestamp with time zone NOT NULL,
	last_contact_failure timestamp with time zone NOT NULL,
	contained boolean NOT NULL,
	disqualified timestamp with time zone,
	PRIMARY KEY ( id )
);
CREATE TABLE offers (
	id serial NOT NULL,
	name text NOT NULL,
	description text NOT NULL,
	award_credit_in_cents integer NOT NULL,
	invitee_credit_in_cents integer NOT NULL,
	award_credit_duration_days integer NOT NULL,
	invitee_credit_duration_days integer NOT NULL,
	redeemable_cap integer NOT NULL,
	num_redeemed integer NOT NULL,
	expires_at timestamp with time zone NOT NULL,
	created_at timestamp with time zone NOT NULL,
	status integer NOT NULL,
	type integer NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE pending_audits (
	node_id bytea NOT NULL,
	piece_id bytea NOT NULL,
	stripe_index bigint NOT NULL,
	share_size bigint NOT NULL,
	expected_share_hash bytea NOT NULL,
	reverify_count bigint NOT NULL,
	PRIMARY KEY ( node_id )
);
CREATE TABLE projects (
	id bytea NOT NULL,
	name text NOT NULL,
	description text NOT NULL,
	usage_limit bigint NOT NULL,
	created_at timestamp with time zone NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE registration_tokens (
	secret bytea NOT NULL,
	owner_id bytea,
	project_limit integer NOT NULL,
	created_at timestamp with time zone NOT NULL,
	PRIMARY KEY ( secret ),
	UNIQUE ( owner_id )
);
CREATE TABLE reset_password_tokens (
	secret bytea NOT NULL,
	owner_id bytea NOT NULL,
	created_at timestamp with time zone NOT NULL,
	PRIMARY KEY ( secret ),
	UNIQUE ( owner_id )
);
CREATE TABLE serial_numbers (
	id serial NOT NULL,
	serial_number bytea NOT NULL,
	bucket_id bytea NOT NULL,
	expires_at timestamp NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE storagenode_bandwidth_rollups (
	storagenode_id bytea NOT NULL,
	interval_start timestamp NOT NULL,
	interval_seconds integer NOT NULL,
	action integer NOT NULL,
	allocated bigint NOT NULL,
	settled bigint NOT NULL,
	PRIMARY KEY ( storagenode_id, interval_start, action )
);
CREATE TABLE storagenode_storage_tallies (
	id bigserial NOT NULL,
	node_id bytea NOT NULL,
	interval_end_time timestamp with time zone NOT NULL,
	data_total double precision NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE users (
	id bytea NOT NULL,
	full_name text NOT NULL,
	short_name text,
	email text NOT NULL,
	password_hash bytea NOT NULL,
	status integer NOT NULL,
	created_at timestamp with time zone NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE api_keys (
	id bytea NOT NULL,
	project_id bytea NOT NULL REFERENCES projects( id ) ON DELETE CASCADE,
	head bytea NOT NULL,
	name text NOT NULL,
	secret bytea NOT NULL,
	created_at timestamp with time zone NOT NULL,
	PRIMARY KEY ( id ),
	UNIQUE ( head ),
	UNIQUE ( name, project_id )
);
CREATE TABLE project_members (
	member_id bytea NOT NULL REFERENCES users( id ) ON DELETE CASCADE,
	project_id bytea NOT NULL REFERENCES projects( id ) ON DELETE CASCADE,
	created_at timestamp with time zone NOT NULL,
	PRIMARY KEY ( member_id, project_id )
);
CREATE TABLE used_serials (
	serial_number_id integer NOT NULL REFERENCES serial_numbers( id ) ON DELETE CASCADE,
	storage_node_id bytea NOT NULL,
	PRIMARY KEY ( serial_number_id, storage_node_id )
);
CREATE TABLE user_credits (
	id serial NOT NULL,
	user_id bytea NOT NULL REFERENCES users( id ),
	offer_id integer NOT NULL REFERENCES offers( id ),
	referred_by bytea REFERENCES users( id ),
	credits_earned_in_cents integer NOT NULL,
	credits_used_in_cents integer NOT NULL,
	expires_at timestamp with time zone NOT NULL,
	created_at timestamp with time zone NOT NULL,
	PRIMARY KEY ( id )
);
CREATE TABLE user_payments (
  user_id bytea NOT NULL REFERENCES users( id ) ON DELETE CASCADE,
  customer_id bytea NOT NULL,
  created_at timestamp with time zone NOT NULL,
  PRIMARY KEY ( user_id ),
  UNIQUE ( customer_id )
);
CREATE TABLE project_payments (
  project_id bytea NOT NULL REFERENCES projects( id ) ON DELETE CASCADE,
  payer_id bytea NOT NULL REFERENCES user_payments( user_id ) ON DELETE CASCADE,
  payment_method_id bytea NOT NULL,
  created_at timestamp with time zone NOT NULL,
  PRIMARY KEY ( project_id )
);
CREATE TABLE project_invoice_stamps (
  project_id bytea NOT NULL REFERENCES projects( id ) ON DELETE CASCADE,
  invoice_id bytea NOT NULL,
  start_date timestamp with time zone NOT NULL,
  end_date timestamp with time zone NOT NULL,
  created_at timestamp with time zone NOT NULL,
  PRIMARY KEY ( project_id, start_date, end_date ),
  UNIQUE ( invoice_id )
);
CREATE INDEX bucket_name_project_id_interval_start_interval_seconds ON bucket_bandwidth_rollups ( bucket_name, project_id, interval_start, interval_seconds );
CREATE UNIQUE INDEX bucket_id_rollup ON bucket_usages ( bucket_id, rollup_end_time );
CREATE INDEX node_last_ip ON nodes ( last_ip );
CREATE UNIQUE INDEX serial_number ON serial_numbers ( serial_number );
CREATE INDEX serial_numbers_expires_at_index ON serial_numbers ( expires_at );
CREATE INDEX storagenode_id_interval_start_interval_seconds ON storagenode_bandwidth_rollups ( storagenode_id, interval_start, interval_seconds );
CREATE TABLE value_attributions (
    project_id bytea NOT NULL,
    bucket_name bytea NOT NULL,
    partner_id bytea NOT NULL,
    last_updated timestamp NOT NULL,
    PRIMARY KEY ( project_id, bucket_name )
);

---

INSERT INTO "accounting_rollups"("id", "node_id", "start_time", "put_total", "get_total", "get_audit_total", "get_repair_total", "put_repair_total", "at_rest_total") VALUES (1, E'\\367M\\177\\251]t/\\022\\256\\214\\265\\025\\224\\204:\\217\\212\\0102<\\321\\374\\020&\\271Qc\\325\\261\\354\\246\\233'::bytea, '2019-02-09 00:00:00+00', 1000, 2000, 3000, 4000, 0, 5000);

INSERT INTO "accounting_timestamps" VALUES ('LastAtRestTally', '0001-01-01 00:00:00+00');
INSERT INTO "accounting_timestamps" VALUES ('LastRollup', '0001-01-01 00:00:00+00');
INSERT INTO "accounting_timestamps" VALUES ('LastBandwidthTally', '0001-01-01 00:00:00+00');

INSERT INTO "nodes"("id", "address", "last_ip", "protocol", "type", "email", "wallet", "free_bandwidth", "free_disk", "major", "minor", "patch", "hash", "timestamp", "release","latency_90", "audit_success_count", "total_audit_count", "audit_success_ratio", "uptime_success_count", "total_uptime_count", "uptime_ratio", "created_at", "updated_at", "last_contact_success", "last_contact_failure", "contained", "disqualified") VALUES (E'\\153\\313\\233\\074\\327\\177\\136\\070\\346\\001', '127.0.0.1:55516', '', 0, 4, '', '', -1, -1, 0, 1, 0, '', 'epoch', false, 0, 0, 5, 0, 0, 5, 0, '2019-02-14 08:07:31.028103+00', '2019-02-14 08:07:31.108963+00', 'epoch', 'epoch', false, NULL);
INSERT INTO "nodes"("id", "address", "last_ip", "protocol", "type", "email", "wallet", "free_bandwidth", "free_disk", "major", "minor", "patch", "hash", "timestamp", "release","latency_90", "audit_success_count", "total_audit_count", "audit_success_ratio", "uptime_success_count", "total_uptime_count", "uptime_ratio", "created_at", "updated_at", "last_contact_success", "last_contact_failure", "contained", "disqualified") VALUES (E'\\006\\223\\250R\\221\\005\\365\\377v>0\\266\\365\\216\\255?\\347\\244\\371?2\\264\\262\\230\\007<\\001\\262\\263\\237\\247n', '127.0.0.1:55518', '', 0, 4, '', '', -1, -1, 0, 1, 0, '', 'epoch', false, 0, 0, 0, 1, 3, 3, 1, '2019-02-14 08:07:31.028103+00', '2019-02-14 08:07:31.108963+00', 'epoch', 'epoch', false, NULL);
INSERT INTO "nodes"("id", "address", "last_ip", "protocol", "type", "email", "wallet", "free_bandwidth", "free_disk", "major", "minor", "patch", "hash", "timestamp", "release","latency_90", "audit_success_count", "total_audit_count", "audit_success_ratio", "uptime_success_count", "total_uptime_count", "uptime_ratio", "created_at", "updated_at", "last_contact_success", "last_contact_failure", "contained", "disqualified") VALUES (E'\\363\\342\\363\\371>+F\\256\\263\\300\\273|\\342N\\347\\014', '127.0.0.1:55517', '', 0, 4, '', '', -1, -1, 0, 1, 0, '', 'epoch', false, 0, 0, 0, 1, 0, 0, 1, '2019-02-14 08:07:31.028103+00', '2019-02-14 08:07:31.108963+00', 'epoch', 'epoch', false, now());

INSERT INTO "projects"("id", "name", "description", "usage_limit", "created_at") VALUES (E'\\022\\217/\\014\\376!K\\023\\276\\031\\311}m\\236\\205\\300'::bytea, 'ProjectName', 'projects description', 0, '2019-02-14 08:28:24.254934+00');

INSERT INTO "users"("id", "full_name", "short_name", "email", "password_hash", "status", "created_at") VALUES (E'\\363\\311\\033w\\222\\303Ci\\265\\343U\\303\\312\\204",'::bytea, 'Noahson', 'William', '1email1@ukr.net', E'some_readable_hash'::bytea, 1, '2019-02-14 08:28:24.614594+00');
INSERT INTO "projects"("id", "name", "description", "usage_limit", "created_at") VALUES (E'\\363\\342\\363\\371>+F\\256\\263\\300\\273|\\342N\\347\\014'::bytea, 'projName1', 'Test project 1', 0, '2019-02-14 08:28:24.636949+00');
INSERT INTO "project_members"("member_id", "project_id", "created_at") VALUES (E'\\363\\311\\033w\\222\\303Ci\\265\\343U\\303\\312\\204",'::bytea, E'\\363\\342\\363\\371>+F\\256\\263\\300\\273|\\342N\\347\\014'::bytea, '2019-02-14 08:28:24.677953+00');

INSERT INTO "irreparabledbs" ("segmentpath", "segmentdetail", "pieces_lost_count", "seg_damaged_unix_sec", "repair_attempt_count") VALUES ('\x49616d5365676d656e746b6579696e666f30', '\x49616d5365676d656e7464657461696c696e666f30', 10, 1550159554, 10);

INSERT INTO "injuredsegments" ("path", "data") VALUES ('0', '\x0a0130120100');
INSERT INTO "injuredsegments" ("path", "data") VALUES ('here''s/a/great/path', '\x0a136865726527732f612f67726561742f70617468120a0102030405060708090a');
INSERT INTO "injuredsegments" ("path", "data") VALUES ('yet/another/cool/path', '\x0a157965742f616e6f746865722f636f6f6c2f70617468120a0102030405060708090a');
INSERT INTO "injuredsegments" ("path", "data") VALUES ('so/many/iconic/paths/to/choose/from', '\x0a23736f2f6d616e792f69636f6e69632f70617468732f746f2f63686f6f73652f66726f6d120a0102030405060708090a');

INSERT INTO "certrecords" VALUES (E'0Y0\\023\\006\\007*\\206H\\316=\\002\\001\\006\\010*\\206H\\316=\\003\\001\\007\\003B\\000\\004\\360\\267\\227\\377\\253u\\222\\337Y\\324C:GQ\\010\\277v\\010\\315D\\271\\333\\337.\\203\\023=C\\343\\014T%6\\027\\362?\\214\\326\\017U\\334\\000\\260\\224\\260J\\221\\304\\331F\\304\\221\\236zF,\\325\\326l\\215\\306\\365\\200\\022', E'L\\301|\\200\\247}F|1\\320\\232\\037n\\335\\241\\206\\244\\242\\207\\204.\\253\\357\\326\\352\\033Dt\\202`\\022\\325', '2019-02-14 08:07:31.335028+00');

INSERT INTO "bucket_usages" ("id", "bucket_id", "rollup_end_time", "remote_stored_data", "inline_stored_data", "remote_segments", "inline_segments", "objects", "metadata_size", "repair_egress", "get_egress", "audit_egress") VALUES (E'\\153\\313\\233\\074\\327\\177\\136\\070\\346\\001",'::bytea, E'\\366\\146\\032\\321\\316\\161\\070\\133\\302\\271",'::bytea, '2019-03-06 08:28:24.677953+00', 10, 11, 12, 13, 14, 15, 16, 17, 18);

INSERT INTO "registration_tokens" ("secret", "owner_id", "project_limit", "created_at") VALUES (E'\\070\\127\\144\\013\\332\\344\\102\\376\\306\\056\\303\\130\\106\\132\\321\\276\\321\\274\\170\\264\\054\\333\\221\\116\\154\\221\\335\\070\\220\\146\\344\\216'::bytea, null, 1, '2019-02-14 08:28:24.677953+00');

INSERT INTO "serial_numbers" ("id", "serial_number", "bucket_id", "expires_at") VALUES (1, E'0123456701234567'::bytea, E'\\363\\342\\363\\371>+F\\256\\263\\300\\273|\\342N\\347\\014/testbucket'::bytea, '2019-03-06 08:28:24.677953+00');
INSERT INTO "used_serials" ("serial_number_id", "storage_node_id") VALUES (1, E'\\006\\223\\250R\\221\\005\\365\\377v>0\\266\\365\\216\\255?\\347\\244\\371?2\\264\\262\\230\\007<\\001\\262\\263\\237\\247n');

INSERT INTO "storagenode_bandwidth_rollups" ("storagenode_id", "interval_start", "interval_seconds", "action", "allocated", "settled") VALUES (E'\\006\\223\\250R\\221\\005\\365\\377v>0\\266\\365\\216\\255?\\347\\244\\371?2\\264\\262\\230\\007<\\001\\262\\263\\237\\247n', '2019-03-06 08:00:00.000000+00', 3600, 1, 1024, 2024);
INSERT INTO "storagenode_storage_tallies" VALUES (1, E'\\3510\\323\\225"~\\036<\\342\\330m\\0253Jhr\\246\\233K\\246#\\2303\\351\\256\\275j\\212UM\\362\\207', '2019-02-14 08:16:57.812849+00', 1000);

INSERT INTO "bucket_bandwidth_rollups" ("bucket_name", "project_id", "interval_start", "interval_seconds", "action", "inline", "allocated", "settled") VALUES (E'testbucket'::bytea, E'\\363\\342\\363\\371>+F\\256\\263\\300\\273|\\342N\\347\\014'::bytea,'2019-03-06 08:00:00.000000+00', 3600, 1, 1024, 2024, 3024);
INSERT INTO "bucket_storage_tallies" ("bucket_name", "project_id", "interval_start", "inline", "remote", "remote_segments_count", "inline_segments_count", "object_count", "metadata_size") VALUES (E'testbucket'::bytea, E'\\363\\342\\363\\371>+F\\256\\263\\300\\273|\\342N\\347\\014'::bytea,'2019-03-06 08:00:00.000000+00', 4024, 5024, 0, 0, 0, 0);

INSERT INTO "reset_password_tokens" ("secret", "owner_id", "created_at") VALUES (E'\\070\\127\\144\\013\\332\\344\\102\\376\\306\\056\\303\\130\\106\\132\\321\\276\\321\\274\\170\\264\\054\\333\\221\\116\\154\\221\\335\\070\\220\\146\\344\\216'::bytea, E'\\363\\311\\033w\\222\\303Ci\\265\\343U\\303\\312\\204",'::bytea, '2019-05-08 08:28:24.677953+00');

INSERT INTO "pending_audits" ("node_id", "piece_id", "stripe_index", "share_size", "expected_share_hash", "reverify_count") VALUES (E'\\153\\313\\233\\074\\327\\177\\136\\070\\346\\001'::bytea, E'\\363\\311\\033w\\222\\303Ci\\265\\343U\\303\\312\\204",'::bytea, 5, 1024, E'\\070\\127\\144\\013\\332\\344\\102\\376\\306\\056\\303\\130\\106\\132\\321\\276\\321\\274\\170\\264\\054\\333\\221\\116\\154\\221\\335\\070\\220\\146\\344\\216'::bytea, 1);

INSERT INTO "offers" ("id", "name", "description", "award_credit_in_cents", "invitee_credit_in_cents", "award_credit_duration_days", "invitee_credit_duration_days", "redeemable_cap", "expires_at", "created_at", "num_redeemed", "status", "type") VALUES (1, 'testOffer', 'Test offer 1', 0, 0, 14, 14, 50, '2019-03-14 08:28:24.636949+00', '2019-02-14 08:28:24.636949+00', 0, 0, 0);

INSERT INTO "api_keys" ("id", "project_id", "head", "name", "secret", "created_at") VALUES (E'\\334/\\302;\\225\\355O\\323\\276f\\247\\354/6\\241\\033'::bytea, E'\\022\\217/\\014\\376!K\\023\\276\\031\\311}m\\236\\205\\300'::bytea, E'\\111\\142\\147\\304\\132\\375\\070\\163\\270\\160\\251\\370\\126\\063\\351\\037\\257\\071\\143\\375\\351\\320\\253\\232\\220\\260\\075\\173\\306\\307\\115\\136'::bytea, 'key 2', E'\\254\\011\\315\\333\\273\\365\\001\\071\\024\\154\\253\\332\\301\\216\\361\\074\\221\\367\\251\\231\\274\\333\\300\\367\\001\\272\\327\\111\\315\\123\\042\\016'::bytea, '2019-02-14 08:28:24.267934+00');

INSERT INTO "user_payments" ("user_id", "customer_id", "created_at") VALUES (E'\\363\\311\\033w\\222\\303Ci\\265\\343U\\303\\312\\204",'::bytea, E'\\022\\217/\\014\\376!K\\023\\276'::bytea, '2019-06-01 08:28:24.267934+00');
INSERT INTO "project_payments" ("project_id", "payer_id", "payment_method_id", "created_at") VALUES (E'\\022\\217/\\014\\376!K\\023\\276\\031\\311}m\\236\\205\\300'::bytea, E'\\363\\311\\033w\\222\\303Ci\\265\\343U\\303\\312\\204",'::bytea, E'\\022\\217/\\014\\376!K\\023\\276'::bytea, '2019-06-01 08:28:24.267934+00');
INSERT INTO "project_invoice_stamps" ("project_id", "invoice_id", "start_date", "end_date", "created_at") VALUES (E'\\022\\217/\\014\\376!K\\023\\276\\031\\311}m\\236\\205\\300'::bytea, E'\\363\\311\\033w\\222\\303,'::bytea, '2019-06-01 08:28:24.267934+00', '2019-06-29 08:28:24.267934+00', '2019-06-01 08:28:24.267934+00');

INSERT INTO "value_attributions" ("project_id", "bucket_name", "partner_id", "last_updated") VALUES (E'\\363\\311\\033w\\222\\303Ci\\265\\343U\\303\\312\\204",'::bytea, E''::bytea, E'\\363\\342\\363\\371>+F\\256\\263\\300\\273|\\342N\\347\\014'::bytea,'2019-02-14 08:07:31.028103+00');

-- NEW DATA --

INSERT INTO "user_credits" ("id", "user_id", "offer_id", "referred_by", "credits_earned_in_cents", "credits_used_in_cents", "expires_at", "created_at") VALUES (1, E'\\363\\311\\033w\\222\\303Ci\\265\\343U\\303\\312\\204",'::bytea, 1, E'\\363\\311\\033w\\222\\303Ci\\265\\343U\\303\\312\\204",'::bytea, 200, 0, '2019-10-01 08:28:24.267934+00', '2019-06-01 08:28:24.267934+00');
