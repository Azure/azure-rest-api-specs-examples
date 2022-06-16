const { MariaDBManagementClient } = require("@azure/arm-mariadb");
const { DefaultAzureCredential } = require("@azure/identity");

async function recommendedActionSessionResult() {
  const subscriptionId = "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const locationName = "WestUS";
  const operationId = "aaaabbbb-cccc-dddd-0000-111122223333";
  const credential = new DefaultAzureCredential();
  const client = new MariaDBManagementClient(credential, subscriptionId);
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
