// Code generated by protoc-gen-grpc-js-client
// source: master.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function MasterCreate(p, conf) {
	path = '/ci/v1beta1/masters'
	return xhr(path, 'POST', conf, null, p);
}

function MasterDelete(p, conf) {
	path = '/ci/v1beta1/masters'
	return xhr(path, 'DELETE', conf, p);
}

module.exports = {
  Master: {
      Create: MasterCreate,
      Delete: MasterDelete
  }
};
