const { AzureTrafficCollectorClient } = require("@azure/arm-networkfunction");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the collector policy in a specified Traffic Collector
 *
 * @summary Gets the collector policy in a specified Traffic Collector
 * x-ms-original-file: specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-11-01/examples/CollectorPolicyGet.json
 */
async function getCollectionPolicy() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const azureTrafficCollectorName = "atc";
  const collectorPolicyName = "cp1";
  const credential = new DefaultAzureCredential();
  const client = new AzureTrafficCollectorClient(credential, subscriptionId);
  const result = await client.collectorPolicies.get(
    resourceGroupName,
    azureTrafficCollectorName,
    collectorPolicyName
  );
  console.log(result);
}

getCollectionPolicy().catch(console.error);
