const { AppConfigurationManagementClient } = require("@azure/arm-appconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a configuration store with the specified parameters.
 *
 * @summary Creates a configuration store with the specified parameters.
 * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2022-05-01/examples/ConfigurationStoresCreate.json
 */
async function configurationStoresCreate() {
  const subscriptionId = "c80fb759-c965-4c6a-9110-9b2b2d038882";
  const resourceGroupName = "myResourceGroup";
  const configStoreName = "contoso";
  const configStoreCreationParameters = {
    location: "westus",
    sku: { name: "Standard" },
    tags: { myTag: "myTagValue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppConfigurationManagementClient(credential, subscriptionId);
  const result = await client.configurationStores.beginCreateAndWait(
    resourceGroupName,
    configStoreName,
    configStoreCreationParameters
  );
  console.log(result);
}

configurationStoresCreate().catch(console.error);
