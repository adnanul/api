package v1beta1

// Auto-generated. DO NOT EDIT.
import (
	"github.com/golang/glog"
	"github.com/xeipuuv/gojsonschema"
)

var clusterInstanceByIPRequestSchema *gojsonschema.Schema
var clusterUpdateRequestSchema *gojsonschema.Schema
var clusterDeleteRequestSchema *gojsonschema.Schema
var clusterCreateRequestSchema *gojsonschema.Schema
var clusterDescribeRequestSchema *gojsonschema.Schema
var clusterListRequestSchema *gojsonschema.Schema
var clusterReconfigureRequestSchema *gojsonschema.Schema
var clusterStartupConfigRequestSchema *gojsonschema.Schema
var clusterClientConfigRequestSchema *gojsonschema.Schema

func init() {
	var err error
	clusterInstanceByIPRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "externalIP": {
      "type": "string"
    },
    "uid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	clusterUpdateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "v1beta1ClusterSettings": {
      "properties": {
        "logIndexPrefix": {
          "type": "string"
        },
        "logStorageLifetime": {
          "title": "Number of secs logs will be stored in ElasticSearch",
          "type": "integer"
        },
        "monitoringStorageLifetime": {
          "title": "Number of secs logs will be stored in InfluxDB",
          "type": "integer"
        }
      },
      "type": "object"
    }
  },
  "properties": {
    "defaultAccessLevel": {
      "title": "Default access level is to allow permission to the cluster\nwhen no Role matched for that specif user or group. This can\nset as\n  - kubernetes:team-admin\n  - kubernetes:cluster-admin\n  - kubernetes:admin\n  - kubernetes:editor\n  - kubernetes:viewer\n  - deny-access",
      "type": "string"
    },
    "doNotDelete": {
      "type": "boolean"
    },
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "settings": {
      "$ref": "#/definitions/v1beta1ClusterSettings"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	clusterDeleteRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "deleteDynamicVolumes": {
      "type": "boolean"
    },
    "force": {
      "type": "boolean"
    },
    "keepLodabalancers": {
      "type": "boolean"
    },
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "releaseReservedIP": {
      "type": "boolean"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	clusterCreateRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "definitions": {
    "v1beta1InstanceGroup": {
      "properties": {
        "count": {
          "type": "integer"
        },
        "sku": {
          "type": "string"
        },
        "useSpotInstances": {
          "type": "boolean"
        }
      },
      "type": "object"
    }
  },
  "properties": {
    "credentialUID": {
      "type": "string"
    },
    "defaultAccessLevel": {
      "title": "Default access level is to allow permission to the cluster\nwhen no Role matched for that specif user or group. This can\nset as\n  - kubernetes:team-admin\n  - kubernetes:cluster-admin\n  - kubernetes:admin\n  - kubernetes:editor\n  - kubernetes:viewer\n  - deny-access\nIf not set this will set \"\"",
      "type": "string"
    },
    "doNotDelete": {
      "type": "boolean"
    },
    "gceProject": {
      "type": "string"
    },
    "kubernetesVersion": {
      "type": "string"
    },
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "nodeGroups": {
      "items": {
        "$ref": "#/definitions/v1beta1InstanceGroup"
      },
      "type": "array"
    },
    "provider": {
      "type": "string"
    },
    "zone": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	clusterDescribeRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "uid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	clusterListRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "status": {
      "items": {
        "type": "string"
      },
      "type": "array"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	clusterReconfigureRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "applyToMaster": {
      "type": "boolean"
    },
    "count": {
      "type": "integer"
    },
    "hostfactsVersion": {
      "type": "string"
    },
    "kubeStarterVersion": {
      "type": "string"
    },
    "kubeletVersion": {
      "type": "string"
    },
    "name": {
      "maxLength": 63,
      "pattern": "^[a-z0-9](?:[a-z0-9\\-]{0,61}[a-z0-9])?$",
      "type": "string"
    },
    "saltbaseVersion": {
      "type": "string"
    },
    "sku": {
      "type": "string"
    },
    "version": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	clusterStartupConfigRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "resourceVersion": {
      "type": "integer"
    },
    "role": {
      "type": "string"
    },
    "uid": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
	clusterClientConfigRequestSchema, err = gojsonschema.NewSchema(gojsonschema.NewStringLoader(`{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "properties": {
    "name": {
      "type": "string"
    }
  },
  "type": "object"
}`))
	if err != nil {
		glog.Fatal(err)
	}
}

func (m *ClusterInstanceByIPRequest) IsValid() (*gojsonschema.Result, error) {
	return clusterInstanceByIPRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterInstanceByIPRequest) IsRequest() {}

func (m *ClusterUpdateRequest) IsValid() (*gojsonschema.Result, error) {
	return clusterUpdateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterUpdateRequest) IsRequest() {}

func (m *ClusterDeleteRequest) IsValid() (*gojsonschema.Result, error) {
	return clusterDeleteRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterDeleteRequest) IsRequest() {}

func (m *ClusterCreateRequest) IsValid() (*gojsonschema.Result, error) {
	return clusterCreateRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterCreateRequest) IsRequest() {}

func (m *ClusterDescribeRequest) IsValid() (*gojsonschema.Result, error) {
	return clusterDescribeRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterDescribeRequest) IsRequest() {}

func (m *ClusterListRequest) IsValid() (*gojsonschema.Result, error) {
	return clusterListRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterListRequest) IsRequest() {}

func (m *ClusterReconfigureRequest) IsValid() (*gojsonschema.Result, error) {
	return clusterReconfigureRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterReconfigureRequest) IsRequest() {}

func (m *ClusterStartupConfigRequest) IsValid() (*gojsonschema.Result, error) {
	return clusterStartupConfigRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterStartupConfigRequest) IsRequest() {}

func (m *ClusterClientConfigRequest) IsValid() (*gojsonschema.Result, error) {
	return clusterClientConfigRequestSchema.Validate(gojsonschema.NewGoLoader(m))
}
func (m *ClusterClientConfigRequest) IsRequest() {}

