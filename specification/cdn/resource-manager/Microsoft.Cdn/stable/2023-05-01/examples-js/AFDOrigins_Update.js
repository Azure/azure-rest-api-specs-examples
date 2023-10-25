const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing origin within an origin group.
 *
 * @summary Updates an existing origin within an origin group.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2023-05-01/examples/AFDOrigins_Update.json
 */
async function afdOriginsUpdate() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "RG";
  const profileName = "profile1";
  const originGroupName = "origingroup1";
  const originName = "origin1";
  const originUpdateProperties = {
    enabledState: "Enabled",
    hostName: "host1.blob.core.windows.net",
    httpPort: 80,
    httpsPort: 443,
  };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.afdOrigins.beginUpdateAndWait(
    resourceGroupName,
    profileName,
    originGroupName,
    originName,
    originUpdateProperties
  );
  console.log(result);
}
