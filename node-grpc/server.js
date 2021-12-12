const grpc = require("@grpc/grpc-js");
const PROTO_PATH = "./user.proto";
var protoLoader = require("@grpc/proto-loader");

const options = {
    keepCase: true,
    longs: String,
    enums: String,
    defaults: true,
    oneofs: true,
};
var packageDefinition = protoLoader.loadSync(PROTO_PATH, options);
const userProto = grpc.loadPackageDefinition(packageDefinition).user;

const server = new grpc.Server();

let user = [
  {
    id: 1,
    username: 'First1',
    email: 'abcd@abcd.com',
    password: 'Last1',
  },
  {
    id: 2,
    username: 'First2',
    email: 'xyz@xyz.com',
    password: 'Last2',
  }
];
server.addService(userProto.UserService.service, {
    ReadAll: (_, callback) => {
      callback(null, user);
    },
  });
  
  server.bindAsync(
    "127.0.0.1:50051",
    grpc.ServerCredentials.createInsecure(),
    (error, port) => {
      console.log("Server running at http://127.0.0.1:50051");
      server.start();
    }
  );