const { SqlManagementClient } = require("@azure/arm-sql");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the failover groups in a location.
 *
 * @summary Lists the failover groups in a location.
 * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/InstanceFailoverGroupList.json
 */
async function listFailoverGroup() {
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = "Default";
  const locationName = "Japan East";
  const credential = new DefaultAzureCredential();
  const client = new SqlManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.instanceFailoverGroups.listByLocation(
    resourceGroupName,
    locationName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listFailoverGroup().catch(console.error);
