const { HybridNetworkManagementClient } = require("@azure/arm-hybridnetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a hybrid configuration group tags.
 *
 * @summary Updates a hybrid configuration group tags.
 * x-ms-original-file: specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/ConfigurationGroupValueUpdateTags.json
 */
async function updateHybridConfigurationGroupTags() {
  const subscriptionId = process.env["HYBRIDNETWORK_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["HYBRIDNETWORK_RESOURCE_GROUP"] || "rg1";
  const configurationGroupValueName = "testConfigurationGroupValue";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new HybridNetworkManagementClient(credential, subscriptionId);
  const result = await client.configurationGroupValues.updateTags(
    resourceGroupName,
    configurationGroupValueName,
    parameters,
  );
  console.log(result);
}
