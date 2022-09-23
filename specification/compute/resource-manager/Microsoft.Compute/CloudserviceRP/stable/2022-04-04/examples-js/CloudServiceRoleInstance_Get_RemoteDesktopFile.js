const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a remote desktop file for a role instance in a cloud service.
 *
 * @summary Gets a remote desktop file for a role instance in a cloud service.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-04-04/examples/CloudServiceRoleInstance_Get_RemoteDesktopFile.json
 */
async function getCloudServiceRole() {
  const subscriptionId = "{subscription-id}";
  const roleInstanceName = "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa";
  const resourceGroupName = "rgcloudService";
  const cloudServiceName = "aaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.cloudServiceRoleInstances.getRemoteDesktopFile(
    roleInstanceName,
    resourceGroupName,
    cloudServiceName
  );
  console.log(result);
}

getCloudServiceRole().catch(console.error);
