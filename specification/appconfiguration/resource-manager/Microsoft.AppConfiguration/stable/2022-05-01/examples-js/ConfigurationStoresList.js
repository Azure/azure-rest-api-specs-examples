const { AppConfigurationManagementClient } = require("@azure/arm-appconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the configuration stores for a given subscription.
 *
 * @summary Lists the configuration stores for a given subscription.
 * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2022-05-01/examples/ConfigurationStoresList.json
 */
async function configurationStoresList() {
  const subscriptionId = "c80fb759-c965-4c6a-9110-9b2b2d038882";
  const credential = new DefaultAzureCredential();
  const client = new AppConfigurationManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.configurationStores.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

configurationStoresList().catch(console.error);
