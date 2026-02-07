const { AppConfigurationManagementClient } = require("@azure/arm-appconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates a configuration store with the specified parameters.
 *
 * @summary updates a configuration store with the specified parameters.
 * x-ms-original-file: 2025-06-01-preview/ConfigurationStoresUpdate.json
 */
async function configurationStoresUpdate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "c80fb759-c965-4c6a-9110-9b2b2d038882";
  const client = new AppConfigurationManagementClient(credential, subscriptionId);
  const result = await client.configurationStores.update("myResourceGroup", "contoso", {
    sku: { name: "Standard" },
    tags: { Category: "Marketing" },
  });
  console.log(result);
}
