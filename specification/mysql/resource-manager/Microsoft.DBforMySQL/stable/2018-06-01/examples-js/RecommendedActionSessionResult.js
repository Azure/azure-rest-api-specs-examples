const { MySQLManagementClient } = require("@azure/arm-mysql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Recommendation action session operation result.
 *
 * @summary Recommendation action session operation result.
 * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/RecommendedActionSessionResult.json
 */
async function recommendedActionSessionResult() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const locationName = "WestUS";
  const operationId = "aaaabbbb-cccc-dddd-0000-111122223333";
  const credential = new DefaultAzureCredential();
  const client = new MySQLManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.locationBasedRecommendedActionSessionsResult.list(
    locationName,
    operationId
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

recommendedActionSessionResult().catch(console.error);
