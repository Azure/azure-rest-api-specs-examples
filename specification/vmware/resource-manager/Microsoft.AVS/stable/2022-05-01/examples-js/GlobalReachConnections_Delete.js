const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a global reach connection in a private cloud
 *
 * @summary Delete a global reach connection in a private cloud
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2022-05-01/examples/GlobalReachConnections_Delete.json
 */
async function globalReachConnectionsDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "group1";
  const privateCloudName = "cloud1";
  const globalReachConnectionName = "connection1";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.globalReachConnections.beginDeleteAndWait(
    resourceGroupName,
    privateCloudName,
    globalReachConnectionName
  );
  console.log(result);
}

globalReachConnectionsDelete().catch(console.error);
