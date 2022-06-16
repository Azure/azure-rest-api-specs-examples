const { AppConfigurationManagementClient } = require("@azure/arm-appconfiguration");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the deleted configuration stores in a subscription.
 *
 * @summary Gets information about the deleted configuration stores in a subscription.
 * x-ms-original-file: specification/appconfiguration/resource-manager/Microsoft.AppConfiguration/stable/2022-05-01/examples/DeletedConfigurationStoresList.json
 */
async function deletedConfigurationStoresList() {
  const subscriptionId = "c80fb759-c965-4c6a-9110-9b2b2d038882";
  const credential = new DefaultAzureCredential();
  const client = new AppConfigurationManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.configurationStores.listDeleted()) {
    resArray.push(item);
  }
  console.log(resArray);
}

deletedConfigurationStoresList().catch(console.error);
