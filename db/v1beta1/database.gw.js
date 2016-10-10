// Code generated by protoc-gen-grpc-js-client
// source: database.proto
// DO NOT EDIT!

/*
This is a RSVP based Ajax client for gRPC gateway JSON APIs.
*/

var xhr = require('grpc-xhr');

function DatabasesList(p, conf) {
	path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/databases'
	delete p['cluster']
	return xhr(path, 'GET', conf, p);
}

function DatabasesCreate(p, conf) {
	path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/databases'
	delete p['cluster']
	return xhr(path, 'POST', conf, null, p);
}

function DatabasesScale(p, conf) {
	path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/databases/' + p['uid'] + '/actions/scale'
	delete p['cluster']
	delete p['uid']
	return xhr(path, 'PUT', conf, null, p);
}

function DatabasesUpdate(p, conf) {
	path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/databases/' + p['uid']
	delete p['cluster']
	delete p['uid']
	return xhr(path, 'PUT', conf, null, p);
}

function DatabasesDescribe(p, conf) {
	path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/databases/' + p['uid']
	delete p['cluster']
	delete p['uid']
	return xhr(path, 'GET', conf, p);
}

function DatabasesDelete(p, conf) {
	path = '/kubernetes/v1beta1/clusters/' + p['cluster'] + '/databases/' + p['uid']
	delete p['cluster']
	delete p['uid']
	return xhr(path, 'DELETE', conf, p);
}

module.exports = {
  Databases: {
      List: DatabasesList,
      Create: DatabasesCreate,
      Scale: DatabasesScale,
      Update: DatabasesUpdate,
      Describe: DatabasesDescribe,
      Delete: DatabasesDelete
  }
};
