// Code generated by protoc-gen-grpc-js-client
// source: quota.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function quotasGet(p, conf) {
    path = '/billing/v1beta1/quotas'
    return xhr(path, 'GET', conf, p);
}

module.exports = {
    quotas: {
        get: quotasGet
    }
};