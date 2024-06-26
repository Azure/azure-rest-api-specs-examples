const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Starts the cloud service.
 *
 * @summary Starts the cloud service.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/CloudserviceRP/stable/2022-04-04/examples/CloudService_Start.json
 */
async function startCloudService() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "ConstosoRG";
  const cloudServiceName = "{cs-name}";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.cloudServices.beginStartAndWait(resourceGroupName, cloudServiceName);
  console.log(result);
}

startCloudService().catch(console.error);
