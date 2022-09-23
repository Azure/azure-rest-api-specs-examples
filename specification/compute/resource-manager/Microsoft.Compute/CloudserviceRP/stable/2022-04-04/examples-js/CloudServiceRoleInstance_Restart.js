const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The Reboot Role Instance asynchronous operation requests a reboot of a role instance in the cloud service.
 *
 * @summary The Reboot Role Instance asynchronous operation requests a reboot of a role instance in the cloud service.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-04-04/examples/CloudServiceRoleInstance_Restart.json
 */
async function restartCloudServiceRoleInstance() {
  const subscriptionId = "{subscription-id}";
  const roleInstanceName = "{roleInstance-name}";
  const resourceGroupName = "ConstosoRG";
  const cloudServiceName = "{cs-name}";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.cloudServiceRoleInstances.beginRestartAndWait(
    roleInstanceName,
    resourceGroupName,
    cloudServiceName
  );
  console.log(result);
}

restartCloudServiceRoleInstance().catch(console.error);
