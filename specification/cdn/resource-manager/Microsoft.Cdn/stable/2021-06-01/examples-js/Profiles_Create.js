const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new Azure Front Door Standard or Azure Front Door Premium or CDN profile with a profile name under the specified subscription and resource group.
 *
 * @summary Creates a new Azure Front Door Standard or Azure Front Door Premium or CDN profile with a profile name under the specified subscription and resource group.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Profiles_Create.json
 */
async function profilesCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const profile = {
    location: "WestCentralUs",
    sku: { name: "Premium_AzureFrontDoor" },
  };
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.profiles.beginCreateAndWait(resourceGroupName, profileName, profile);
  console.log(result);
}

profilesCreate().catch(console.error);
