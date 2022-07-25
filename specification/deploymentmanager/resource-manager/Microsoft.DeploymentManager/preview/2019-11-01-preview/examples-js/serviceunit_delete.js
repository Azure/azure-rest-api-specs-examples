const { AzureDeploymentManager } = require("@azure/arm-deploymentmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the service unit.
 *
 * @summary Deletes the service unit.
 * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/serviceunit_delete.json
 */
async function deleteServiceUnit() {
  const subscriptionId = "caac1590-e859-444f-a9e0-62091c0f5929";
  const resourceGroupName = "myResourceGroup";
  const serviceTopologyName = "myTopology";
  const serviceName = "myService";
  const serviceUnitName = "myServiceUnit";
  const credential = new DefaultAzureCredential();
  const client = new AzureDeploymentManager(credential, subscriptionId);
  const result = await client.serviceUnits.delete(
    resourceGroupName,
    serviceTopologyName,
    serviceName,
    serviceUnitName
  );
  console.log(result);
}

deleteServiceUnit().catch(console.error);
