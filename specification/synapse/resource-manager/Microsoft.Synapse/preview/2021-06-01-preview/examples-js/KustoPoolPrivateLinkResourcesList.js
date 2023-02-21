const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all Kusto pool PrivateLinkResources.
 *
 * @summary Lists all Kusto pool PrivateLinkResources.
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolPrivateLinkResourcesList.json
 */
async function kustoPoolPrivateLinkResourcesList() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "7a587823-959d-4ad0-85bd-cf2a7cef436a";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "DP-900";
  const workspaceName = "synapse-ws-ebi-data";
  const kustoPoolName = "dataexplorerpool900";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.kustoPoolPrivateLinkResourcesOperations.list(
    resourceGroupName,
    workspaceName,
    kustoPoolName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
