const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Deletes an routing rule collection.
 *
 * @summary Deletes an routing rule collection.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2024-07-01/examples/NetworkManagerRoutingRuleCollectionDelete.json
 */
async function deletesAnRoutingRuleCollection() {
  const subscriptionId =
    process.env["NETWORK_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["NETWORK_RESOURCE_GROUP"] || "rg1";
  const networkManagerName = "testNetworkManager";
  const configurationName = "myTestRoutingConfig";
  const ruleCollectionName = "testRuleCollection";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.routingRuleCollections.beginDeleteAndWait(
    resourceGroupName,
    networkManagerName,
    configurationName,
    ruleCollectionName,
  );
  console.log(result);
}
