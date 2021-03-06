const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a role from a cloud service.
 *
 * @summary Gets a role from a cloud service.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/GetCloudServiceRole.json
 */
async function getCloudServiceRole() {
  const subscriptionId = "{subscription-id}";
  const roleName = "{role-name}";
  const resourceGroupName = "ConstosoRG";
  const cloudServiceName = "{cs-name}";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.cloudServiceRoles.get(roleName, resourceGroupName, cloudServiceName);
  console.log(result);
}

getCloudServiceRole().catch(console.error);
