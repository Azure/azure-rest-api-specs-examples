const { MicrosoftDatadogClient } = require("@azure/arm-datadog");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List the api keys for a given monitor resource.
 *
 * @summary List the api keys for a given monitor resource.
 * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2023-01-01/examples/ApiKeys_List.json
 */
async function monitorsListApiKeys() {
  const subscriptionId =
    process.env["DATADOG_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["DATADOG_RESOURCE_GROUP"] || "myResourceGroup";
  const monitorName = "myMonitor";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftDatadogClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.monitors.listApiKeys(resourceGroupName, monitorName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
