const { SpatioClient } = require("@azure/arm-planetarycomputer");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to update a GeoCatalog
 *
 * @summary update a GeoCatalog
 * x-ms-original-file: 2025-02-11-preview/GeoCatalogs_Update.json
 */
async function geoCatalogsUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "cd9b6cdf-dcf0-4dca-ab19-82be07b74704";
  const client = new SpatioClient(credential, subscriptionId);
  const result = await client.geoCatalogs.update("MyResourceGroup", "MyCatalog", {
    tags: { MyTag: "MyValue" },
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/cd9b6cdf-dcf0-4dca-ab19-82be07b74704/resourceGroups/MyResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/MyManagedIdentity":
          {},
      },
    },
  });
  console.log(result);
}
