const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Metadata.
 *
 * @summary Get a Metadata.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/metadata/GetMetadata.json
 */
async function getSingleMetadataByName() {
  const subscriptionId = "2e1dc338-d04d-4443-b721-037eff4fdcac";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const metadataName = "metadataName";
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.metadata.get(resourceGroupName, workspaceName, metadataName);
  console.log(result);
}

getSingleMetadataByName().catch(console.error);
