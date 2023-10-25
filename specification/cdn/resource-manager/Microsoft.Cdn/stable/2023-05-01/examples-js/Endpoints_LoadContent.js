const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Pre-loads a content to CDN. Available for Verizon Profiles.
 *
 * @summary Pre-loads a content to CDN. Available for Verizon Profiles.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/Endpoints_LoadContent.json
 */
async function endpointsLoadContent() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "RG";
  const profileName = "profile1";
  const endpointName = "endpoint1";
  const contentFilePaths = { contentPaths: ["/folder1"] };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.endpoints.beginLoadContentAndWait(
    resourceGroupName,
    profileName,
    endpointName,
    contentFilePaths
  );
  console.log(result);
}
