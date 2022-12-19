const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get all private link resources for a private link hub
 *
 * @summary Get all private link resources for a private link hub
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ListPrivateLinkHubPrivateLinkResources.json
 */
async function getPrivateLinkResourcesForPrivateLinkHub() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "01234567-89ab-4def-0123-456789abcdef";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "ExampleResourceGroup";
  const privateLinkHubName = "ExamplePrivateLinkHub";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateLinkHubPrivateLinkResources.list(
    resourceGroupName,
    privateLinkHubName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
