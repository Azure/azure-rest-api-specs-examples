const { AppConfigurationManagementClient } = require("@azure/arm-appconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a configuration store with the specified parameters.
 *
 * @summary Updates a configuration store with the specified parameters.
 * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2022-05-01/examples/ConfigurationStoresUpdate.json
 */
async function configurationStoresUpdate() {
  const subscriptionId = "c80fb759-c965-4c6a-9110-9b2b2d038882";
  const resourceGroupName = "myResourceGroup";
  const configStoreName = "contoso";
  const configStoreUpdateParameters = {
    sku: { name: "Standard" },
    tags: { category: "Marketing" },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppConfigurationManagementClient(credential, subscriptionId);
  const result = await client.configurationStores.beginUpdateAndWait(
    resourceGroupName,
    configStoreName,
    configStoreUpdateParameters
  );
  console.log(result);
}

configurationStoresUpdate().catch(console.error);
