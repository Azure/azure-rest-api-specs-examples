const { AzureDeploymentManager } = require("@azure/arm-deploymentmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This is an asynchronous operation and can be polled to completion using the operation resource returned by this operation.
 *
 * @summary This is an asynchronous operation and can be polled to completion using the operation resource returned by this operation.
 * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/serviceunit_createorupdate.json
 */
async function createServiceUnitUsingRelativePathsIntoTheArtifactSource() {
  const subscriptionId = "caac1590-e859-444f-a9e0-62091c0f5929";
  const resourceGroupName = "myResourceGroup";
  const serviceTopologyName = "myTopology";
  const serviceName = "myService";
  const serviceUnitName = "myServiceUnit";
  const serviceUnitInfo = {
    artifacts: {
      parametersArtifactSourceRelativePath: "parameter/myTopologyUnit.parameters.json",
      templateArtifactSourceRelativePath: "templates/myTopologyUnit.template.json",
    },
    deploymentMode: "Incremental",
    location: "centralus",
    tags: {},
    targetResourceGroup: "myDeploymentResourceGroup",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureDeploymentManager(credential, subscriptionId);
  const result = await client.serviceUnits.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceTopologyName,
    serviceName,
    serviceUnitName,
    serviceUnitInfo
  );
  console.log(result);
}

createServiceUnitUsingRelativePathsIntoTheArtifactSource().catch(console.error);
