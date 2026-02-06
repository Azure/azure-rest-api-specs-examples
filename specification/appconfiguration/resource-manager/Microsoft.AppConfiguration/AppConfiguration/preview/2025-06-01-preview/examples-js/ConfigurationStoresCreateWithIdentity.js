const { AppConfigurationManagementClient } = require("@azure/arm-appconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a configuration store with the specified parameters.
 *
 * @summary creates a configuration store with the specified parameters.
 * x-ms-original-file: 2025-06-01-preview/ConfigurationStoresCreateWithIdentity.json
 */
async function configurationStoresCreateWithIdentity() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "c80fb759-c965-4c6a-9110-9b2b2d038882";
  const client = new AppConfigurationManagementClient(credential, subscriptionId);
  const result = await client.configurationStores.create("myResourceGroup", "contoso", {
    identity: {
      type: "SystemAssigned, UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourcegroups/myResourceGroup1/providers/Microsoft.ManagedIdentity/userAssignedIdentities/identity2":
          {},
      },
    },
    location: "westus",
    sku: { name: "Standard" },
    tags: { myTag: "myTagValue" },
  });
  console.log(result);
}
