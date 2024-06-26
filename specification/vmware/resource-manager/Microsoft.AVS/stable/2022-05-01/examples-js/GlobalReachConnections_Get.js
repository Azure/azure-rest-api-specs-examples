const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a global reach connection by name in a private cloud
 *
 * @summary Get a global reach connection by name in a private cloud
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/GlobalReachConnections_Get.json
 */
async function globalReachConnectionsGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const globalReachConnectionName = "connection1";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.globalReachConnections.get(
    resourceGroupName,
    privateCloudName,
    globalReachConnectionName
  );
  console.log(result);
}

globalReachConnectionsGet().catch(console.error);
