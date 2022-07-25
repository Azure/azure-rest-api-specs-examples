const { AzureDeploymentManager } = require("@azure/arm-deploymentmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Synchronously creates a new service topology or updates an existing service topology.
 *
 * @summary Synchronously creates a new service topology or updates an existing service topology.
 * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/servicetopology_createorupdate_noartifactsource.json
 */
async function createATopologyWithoutArtifactSource() {
  const subscriptionId = "caac1590-e859-444f-a9e0-62091c0f5929";
  const resourceGroupName = "myResourceGroup";
  const serviceTopologyName = "myTopology";
  const serviceTopologyInfo = {
    location: "centralus",
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureDeploymentManager(credential, subscriptionId);
  const result = await client.serviceTopologies.createOrUpdate(
    resourceGroupName,
    serviceTopologyName,
    serviceTopologyInfo
  );
  console.log(result);
}

createATopologyWithoutArtifactSource().catch(console.error);
