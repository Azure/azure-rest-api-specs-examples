const { AzureDeploymentManager } = require("@azure/arm-deploymentmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to This is an asynchronous operation and can be polled to completion using the operation resource returned by this operation.
 *
 * @summary This is an asynchronous operation and can be polled to completion using the operation resource returned by this operation.
 * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/serviceunit_createorupdate_noartifactsource.json
 */
async function createServiceUnitUsingSasUrIs() {
  const subscriptionId = "caac1590-e859-444f-a9e0-62091c0f5929";
  const resourceGroupName = "myResourceGroup";
  const serviceTopologyName = "myTopology";
  const serviceName = "myService";
  const serviceUnitName = "myServiceUnit";
  const serviceUnitInfo = {
    artifacts: {
      parametersUri:
        "https://mystorageaccount.blob.core.windows.net/myartifactsource/parameter/myTopologyUnit.parameters.json?st=2018-07-07T14%3A10%3A00Z&se=2019-12-31T15%3A10%3A00Z&sp=rl&sv=2017-04-17&sr=c&sig=Yh2SoJ1NhhLRwCLln7de%2Fkabcdefghijklmno5sWEIk%3D",
      templateUri:
        "https://mystorageaccount.blob.core.windows.net/myartifactsource/templates/myTopologyUnit.template.json?st=2018-07-07T14%3A10%3A00Z&se=2019-12-31T15%3A10%3A00Z&sp=rl&sv=2017-04-17&sr=c&sig=Yh2SoJ1NhhLRwCLln7de%2Fkabcdefghijklmno5sWEIk%3D",
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

createServiceUnitUsingSasUrIs().catch(console.error);
