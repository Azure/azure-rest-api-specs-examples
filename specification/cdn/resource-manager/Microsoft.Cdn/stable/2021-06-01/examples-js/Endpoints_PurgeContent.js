const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Removes a content from CDN.
 *
 * @summary Removes a content from CDN.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Endpoints_PurgeContent.json
 */
async function endpointsPurgeContent() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const endpointName = "endpoint1";
  const contentFilePaths = { contentPaths: ["/folder1"] };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.endpoints.beginPurgeContentAndWait(
    resourceGroupName,
    profileName,
    endpointName,
    contentFilePaths
  );
  console.log(result);
}

endpointsPurgeContent().catch(console.error);
