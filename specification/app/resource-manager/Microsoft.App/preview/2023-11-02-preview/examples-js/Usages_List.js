const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets, for the specified location, the current resource usage information as well as the limits under the subscription.
 *
 * @summary Gets, for the specified location, the current resource usage information as well as the limits under the subscription.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/Usages_List.json
 */
async function listUsages() {
  const subscriptionId = process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "subid";
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.usages.list(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
