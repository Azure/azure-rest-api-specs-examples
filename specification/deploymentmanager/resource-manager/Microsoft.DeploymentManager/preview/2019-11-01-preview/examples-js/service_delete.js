const { AzureDeploymentManager } = require("@azure/arm-deploymentmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the service.
 *
 * @summary Deletes the service.
 * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/service_delete.json
 */
async function deleteService() {
  const subscriptionId = "caac1590-e859-444f-a9e0-62091c0f5929";
  const resourceGroupName = "myResourceGroup";
  const serviceTopologyName = "myTopology";
  const serviceName = "myService";
  const credential = new DefaultAzureCredential();
  const client = new AzureDeploymentManager(credential, subscriptionId);
  const result = await client.services.delete(resourceGroupName, serviceTopologyName, serviceName);
  console.log(result);
}

deleteService().catch(console.error);
