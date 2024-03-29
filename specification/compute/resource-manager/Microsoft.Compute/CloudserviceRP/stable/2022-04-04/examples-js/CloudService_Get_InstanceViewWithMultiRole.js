const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the status of a cloud service.
 *
 * @summary Gets the status of a cloud service.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-04-04/examples/CloudService_Get_InstanceViewWithMultiRole.json
 */
async function getCloudServiceInstanceViewWithMultipleRoles() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "ConstosoRG";
  const cloudServiceName = "{cs-name}";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.cloudServices.getInstanceView(resourceGroupName, cloudServiceName);
  console.log(result);
}

getCloudServiceInstanceViewWithMultipleRoles().catch(console.error);
