const { AppConfigurationManagementClient } = require("@azure/arm-appconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a configuration store with the specified parameters.
 *
 * @summary Creates a configuration store with the specified parameters.
 * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2024-05-01/examples/ConfigurationStoresCreateWithDataPlaneProxy.json
 */
async function configurationStoresCreateWithDataPlaneProxy() {
  const subscriptionId =
    process.env["APPCONFIGURATION_SUBSCRIPTION_ID"] || "c80fb759-c965-4c6a-9110-9b2b2d038882";
  const resourceGroupName = process.env["APPCONFIGURATION_RESOURCE_GROUP"] || "myResourceGroup";
  const configStoreName = "contoso";
  const configStoreCreationParameters = {
    dataPlaneProxy: {
      authenticationMode: "Pass-through",
      privateLinkDelegation: "Enabled",
    },
    location: "westus",
    sku: { name: "Standard" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppConfigurationManagementClient(credential, subscriptionId);
  const result = await client.configurationStores.beginCreateAndWait(
    resourceGroupName,
    configStoreName,
    configStoreCreationParameters,
  );
  console.log(result);
}
