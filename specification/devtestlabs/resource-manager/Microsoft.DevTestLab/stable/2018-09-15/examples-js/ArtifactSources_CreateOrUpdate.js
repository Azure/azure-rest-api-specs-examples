const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or replace an existing artifact source.
 *
 * @summary Create or replace an existing artifact source.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/ArtifactSources_CreateOrUpdate.json
 */
async function artifactSourcesCreateOrUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const name = "{artifactSourceName}";
  const artifactSource = {
    armTemplateFolderPath: "{armTemplateFolderPath}",
    branchRef: "{branchRef}",
    displayName: "{displayName}",
    folderPath: "{folderPath}",
    securityToken: "{securityToken}",
    sourceType: "{VsoGit|GitHub|StorageAccount}",
    status: "{Enabled|Disabled}",
    tags: { tagName1: "tagValue1" },
    uri: "{artifactSourceUri}",
  };
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.artifactSources.createOrUpdate(
    resourceGroupName,
    labName,
    name,
    artifactSource
  );
  console.log(result);
}

artifactSourcesCreateOrUpdate().catch(console.error);
