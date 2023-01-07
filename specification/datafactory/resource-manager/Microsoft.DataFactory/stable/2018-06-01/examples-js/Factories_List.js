const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists factories under the specified subscription.
 *
 * @summary Lists factories under the specified subscription.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Factories_List.json
 */
async function factoriesList() {
  const subscriptionId =
    process.env["DATAFACTORY_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-12345678abc";
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.factories.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
