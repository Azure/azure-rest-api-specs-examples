const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Deletes a private endpoint connection under a disk access resource.
 *
 * @summary Deletes a private endpoint connection under a disk access resource.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/diskAccessExamples/DiskAccessPrivateEndpointConnection_Delete.json
 */
async function deleteAPrivateEndpointConnectionUnderADiskAccessResource() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const diskAccessName = "myDiskAccess";
  const privateEndpointConnectionName = "myPrivateEndpointConnection";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.diskAccesses.beginDeleteAPrivateEndpointConnectionAndWait(
    resourceGroupName,
    diskAccessName,
    privateEndpointConnectionName,
  );
  console.log(result);
}
