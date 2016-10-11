// Code generated by protoc-gen-grpc-js-client
// source: conduit.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function conduitWhoAmI(p, conf) {
    path = '/auth/v1beta1/conduit/whoami'
    return xhr(path, 'POST', conf, null, p);
}

function conduitUsers(p, conf) {
    path = '/auth/v1beta1/conduit/users'
    return xhr(path, 'POST', conf, null, p);
}

module.exports = {
    conduit: {
        whoAmI: conduitWhoAmI,
        users: conduitUsers
    }
};