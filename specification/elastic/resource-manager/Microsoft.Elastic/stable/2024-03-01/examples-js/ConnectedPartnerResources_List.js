const { MicrosoftElastic } = require("@azure/arm-elastic");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List of all active deployments that are associated with the marketplace subscription linked to the given monitor.
 *
 * @summary List of all active deployments that are associated with the marketplace subscription linked to the given monitor.
 * x-ms-original-file: specification/elastic/resource-manager/Microsoft.Elastic/stable/2024-03-01/examples/ConnectedPartnerResources_List.json
 */
async function connectedPartnerResourcesList() {
  const subscriptionId =
    process.env["ELASTIC_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["ELASTIC_RESOURCE_GROUP"] || "myResourceGroup";
  const monitorName = "myMonitor";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftElastic(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.connectedPartnerResources.list(resourceGroupName, monitorName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
