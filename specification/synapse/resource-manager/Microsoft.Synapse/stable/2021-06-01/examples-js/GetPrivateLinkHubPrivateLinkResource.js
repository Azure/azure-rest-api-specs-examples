const { SynapseManagementClient } = require("@azure/arm-synapse");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get private link resource in private link hub
 *
 * @summary Get private link resource in private link hub
 * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetPrivateLinkHubPrivateLinkResource.json
 */
async function getPrivateLinkResourcesForPrivateLinkHub() {
  const subscriptionId =
    process.env["SYNAPSE_SUBSCRIPTION_ID"] || "01234567-89ab-4def-0123-456789abcdef";
  const resourceGroupName = process.env["SYNAPSE_RESOURCE_GROUP"] || "ExampleResourceGroup";
  const privateLinkHubName = "ExamplePrivateLinkHub";
  const privateLinkResourceName = "sql";
  const credential = new DefaultAzureCredential();
  const client = new SynapseManagementClient(credential, subscriptionId);
  const result = await client.privateLinkHubPrivateLinkResources.get(
    resourceGroupName,
    privateLinkHubName,
    privateLinkResourceName
  );
  console.log(result);
}
