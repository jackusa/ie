package groups

import (
	"encoding/json"
	"net/http"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	fhir "github.com/intervention-engine/fhir/models"
	"github.com/intervention-engine/fhir/search"
	"github.com/intervention-engine/fhir/server"
	"github.com/intervention-engine/ie/testutil"
	"github.com/stretchr/testify/suite"
	"gopkg.in/mgo.v2/bson"
)

// In order for 'go test' to run this suite, we need to create
// a normal test function and pass our suite to suite.Run
func TestGroupParamSuite(t *testing.T) {
	suite.Run(t, new(GroupParamSuite))
}

type GroupParamSuite struct {
	testutil.MongoSuite
}

func (suite *GroupParamSuite) SetupTest() {
	// Setup the database
	server.Database = suite.DB()
}

func (suite *GroupParamSuite) TearDownTest() {
	suite.TearDownDB()
}

func (suite *GroupParamSuite) TearDownSuite() {
	suite.TearDownDBServer()
}

func (suite *GroupParamSuite) TestGroupParamParser() {
	require := suite.Require()
	assert := suite.Assert()

	p, err := GroupParamParser(GroupParamInfo, search.SearchParamData{Value: "123"})
	require.NoError(err)
	assert.IsType(new(GroupParam), p)
	assert.Equal("123", p.(*GroupParam).Code)
}

func (suite *GroupParamSuite) TestGroupBSONBuilder() {
	require := suite.Require()
	assert := suite.Assert()

	// Store the group
	group := new(fhir.Group)
	suite.InsertFixture("groups", "../fixtures/sample-group.json", group)

	// Store the bundle of data to match the group
	bundleFile, err := os.Open("../fixtures/sample-group-data-bundle.json")
	require.NoError(err)
	ctx, rw, _ := gin.CreateTestContext()
	ctx.Request, err = http.NewRequest("POST", "http://ie-server/", bundleFile)
	require.NoError(err)
	ctx.Request.Header.Add("Content-Type", "application/json")
	server.NewBatchController(server.NewMongoDataAccessLayer(suite.DB())).Post(ctx)
	require.Equal(200, rw.Code)
	bundle := new(fhir.Bundle)
	err = json.NewDecoder(rw.Body).Decode(bundle)
	require.NoError(err)

	// Create the param
	param, err := GroupParamParser(GroupParamInfo, search.SearchParamData{Value: group.Id})
	require.NoError(err)

	obtained, err := GroupBSONBuilder(param, search.NewMongoSearcher(server.Database))
	require.NoError(err)
	expected := bson.M{
		"_id": bson.M{
			"$in": []string{bundle.Entry[0].Resource.(*fhir.Patient).Id},
		},
	}
	assert.Equal(expected, obtained)
}

func (suite *GroupParamSuite) TestQueryParamParser() {
	require := suite.Require()
	assert := suite.Assert()

	p, err := QueryParamParser(QueryParamInfo, search.SearchParamData{Value: "group"})
	require.NoError(err)
	assert.IsType(new(QueryParam), p)
	assert.Equal("group", p.(*QueryParam).String)
}

func (suite *GroupParamSuite) TestQueryBSONBuilder() {
	require := suite.Require()
	assert := suite.Assert()

	// Create the param
	param, err := QueryParamParser(QueryParamInfo, search.SearchParamData{Value: "group"})
	require.NoError(err)

	obtained, err := QueryBSONBuilder(param, search.NewMongoSearcher(server.Database))
	require.NoError(err)
	assert.Equal(bson.M{}, obtained)
}

func (suite *GroupParamSuite) TestQueryBSONBuilderWithWrongQuery() {
	require := suite.Require()
	assert := suite.Assert()

	// Create the param
	param, err := QueryParamParser(QueryParamInfo, search.SearchParamData{Value: "foo"})
	require.NoError(err)

	_, err = QueryBSONBuilder(param, search.NewMongoSearcher(server.Database))
	assert.Error(err)
}
