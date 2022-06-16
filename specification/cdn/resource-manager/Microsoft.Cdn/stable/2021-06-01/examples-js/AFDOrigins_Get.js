const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an existing origin within an origin group.
 *
 * @summary Gets an existing origin within an origin group.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDOrigins_Get.json
 */
async function afdOriginsGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const originGroupName = "origingroup1";
  const originName = "origin1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.afdOrigins.get(
    resourceGroupName,
    profileName,
    originGroupName,
    originName
  );
  console.log(result);
}

afdOriginsGet().catch(console.error);
