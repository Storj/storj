const api = {
  authToken: "",
  operations: {
    user: {
      async create(
        email: string,
        fullName: string,
        password: string
      ): Promise<any> {
        console.log(email, fullName, password);
        return {};
      },
    },
  },
};

export default api;
