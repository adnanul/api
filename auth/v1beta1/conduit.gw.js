// Code generated by protoc-gen-grpc-js-client
// source: conduit.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function ConduitWhoAmI(p, conf) {
	path = '/auth/v1beta1/conduit/whoami'
	return xhr(path, 'POST', conf, null, p);
}

function ConduitUsers(p, conf) {
	path = '/auth/v1beta1/conduit/users'
	return xhr(path, 'POST', conf, null, p);
}

module.exports = {
  Conduit: {
      WhoAmI: ConduitWhoAmI,
      Users: ConduitUsers
  }
};
