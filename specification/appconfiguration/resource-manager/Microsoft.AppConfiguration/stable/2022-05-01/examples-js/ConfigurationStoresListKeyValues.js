const { AppConfigurationManagementClient } = require("@azure/arm-appconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the key-values for a given configuration store.
 *
 * @summary Lists the key-values for a given configuration store.
 * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2022-05-01/examples/ConfigurationStoresListKeyValues.json
 */
async function keyValuesListByConfigurationStore() {
  const subscriptionId = "c80fb759-c965-4c6a-9110-9b2b2d038882";
  const resourceGroupName = "myResourceGroup";
  const configStoreName = "contoso";
  const credential = new DefaultAzureCredential();
  const client = new AppConfigurationManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.keyValues.listByConfigurationStore(
    resourceGroupName,
    configStoreName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

keyValuesListByConfigurationStore().catch(console.error);
