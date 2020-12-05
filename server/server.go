package server

import (
	"bazel-golang-wasm-protoc/protos/api"
	"context"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

////
// ec2Instances file
////

type ec2Instance struct {
	ZoneID          string `json:"zoneID"`
	NetworkTypes    string `json:"networkTypes"`
	RegionID        string `json:"regionID"`
	ZoneStatue      string `json:"zoneStatue"`
	Engine          string `json:"engine"`
	EngineVersion   string `json:"engineVersion"`
	Category        string `json:"category"`
	StorageType     string `json:"storageType"`
	DBInstanceClass string `json:"dbInstanceClass"`
	// DBInstanceRange map[string]interface{} `json:"dbInstanceRange"`
	StorageRange string `json:"storageRange"`
}

////
// Server
////

// func getMSession(conn string) (*mgo.Session, error) {
// 	return mgo.Dial(conn)
// }

var sess *mgo.Session

type Server struct {
	collection *mgo.Collection
	instances  []*api.Instance
}

func init() {
	var err error
	sess, err = mgo.Dial("10.0.1.206")
	if err != nil {
		panic(err)
	}
}

func (server *Server) Search(ctx context.Context, in *api.SearchRequest) (*api.Instances, error) {
	if server.instances == nil {
		server.parseInstances()
	}

	// 主要的逻辑
	instances := []api.Instance{}
	res := []*api.Instance{}

	sc := sess.Clone()
	defer sc.Close()
	var query = bson.M{}
	var count int

	query["dbinstanceclass"] = in.Query
	count, err := sc.DB("alicloud").C("rds").
		Find(query).Count()
	if err != nil {
		panic(err)
	}
	if err := sc.DB("alicloud").C("rds").
		Find(query).
		Skip((int(in.PageIndex) - 1) * int(in.PageSize)).
		Limit(int(in.PageSize)).
		All(&instances); err != nil {
		panic(err)
	}

	for _, inst := range instances {
		res = append(res, &inst)
	}

	return &api.Instances{Instances: res, Count: int32(count)}, nil
}

func (server *Server) parseInstances() {
	// fileName := "external/com_github_ec2instances/file/instances.json"
	ec2Instances := []ec2Instance{}
	server.instances = []*api.Instance{}

	// file, _ := ioutil.ReadFile(fileName)
	// json.Unmarshal(file, &ec2Instances)
	sc := sess.Clone()
	defer sc.Close()

	q := bson.M{}
	err := sc.DB("alicloud").C("rds").Find(q).All(&ec2Instances)
	if err != nil {
		panic(err)
	}

	for _, e := range ec2Instances {
		server.instances = append(server.instances, &api.Instance{
			ZoneID:          e.ZoneID,
			NetworkTypes:    e.NetworkTypes,
			RegionID:        e.RegionID,
			ZoneStatue:      e.ZoneStatue,
			Engine:          e.Engine,
			EngineVersion:   e.EngineVersion,
			Category:        e.Category,
			StorageType:     e.StorageType,
			DBInstanceClass: e.DBInstanceClass,
			// DBInstanceRange: e.DBInstanceRange,
			StorageRange: e.StorageRange,
		})
	}
}
