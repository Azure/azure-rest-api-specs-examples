const { AppConfigurationManagementClient } = require("@azure/arm-appconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a configuration store with the specified parameters.
 *
 * @summary creates a configuration store with the specified parameters.
 * x-ms-original-file: 2025-06-01-preview/ConfigurationStoresCreateWithTelemetry.json
 */
async function configurationStoresCreateWithTelemetry() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "c80fb759-c965-4c6a-9110-9b2b2d038882";
  const client = new AppConfigurationManagementClient(credential, subscriptionId);
  const result = await client.configurationStores.create("myResourceGroup", "contoso", {
    location: "westus",
    telemetry: {
      resourceId:
        "/subscriptions/c80fb759-c965-4c6a-9110-9b2b2d038882/resourceGroups/myResourceGroup/providers/microsoft.insights/components/appInsightsName",
    },
    sku: { name: "Standard" },
    tags: { myTag: "myTagValue" },
  });
  console.log(result);
}
