const { AzureDeploymentManager } = require("@azure/arm-deploymentmanager");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Synchronously creates a new artifact source or updates an existing artifact source.
 *
 * @summary Synchronously creates a new artifact source or updates an existing artifact source.
 * x-ms-original-file: specification/deploymentmanager/resource-manager/Microsoft.DeploymentManager/preview/2019-11-01-preview/examples/artifactsource_createorupdate.json
 */
async function createArtifactSource() {
  const subscriptionId = "caac1590-e859-444f-a9e0-62091c0f5929";
  const resourceGroupName = "myResourceGroup";
  const artifactSourceName = "myArtifactSource";
  const artifactSourceInfo = {
    authentication: {
      type: "Sas",
      sasUri:
        "https://mystorageaccount.blob.core.windows.net/myartifactsource?st=2018-07-07T14%3A10%3A00Z&se=2019-12-31T15%3A10%3A00Z&sp=rl&sv=2017-04-17&sr=c&sig=Yh2SoJ1NhhLRwCLln7de%2Fkabcdefghijklmno5sWEIk%3D",
    },
    location: "centralus",
    sourceType: "AzureStorage",
    tags: {},
  };
  const options = {
    artifactSourceInfo,
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureDeploymentManager(credential, subscriptionId);
  const result = await client.artifactSources.createOrUpdate(
    resourceGroupName,
    artifactSourceName,
    options
  );
  console.log(result);
}

createArtifactSource().catch(console.error);
