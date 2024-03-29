const { HybridNetworkManagementClient } = require("@azure/arm-hybridnetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a network function definition version resource.
 *
 * @summary Updates a network function definition version resource.
 * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/NetworkFunctionDefinitionVersionUpdateTags.json
 */
async function updateTheNetworkFunctionDefinitionVersionTags() {
  const subscriptionId = process.env["HYBRIDNETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HYBRIDNETWORK_RESOURCE_GROUP"] || "rg";
  const publisherName = "TestPublisher";
  const networkFunctionDefinitionGroupName = "TestNetworkFunctionDefinitionGroupName";
  const networkFunctionDefinitionVersionName = "1.0.0";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new HybridNetworkManagementClient(credential, subscriptionId);
  const result = await client.networkFunctionDefinitionVersions.update(
    resourceGroupName,
    publisherName,
    networkFunctionDefinitionGroupName,
    networkFunctionDefinitionVersionName,
    parameters,
  );
  console.log(result);
}
