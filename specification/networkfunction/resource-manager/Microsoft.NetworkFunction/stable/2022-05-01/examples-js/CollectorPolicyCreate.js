const { AzureTrafficCollectorClient } = require("@azure/arm-networkfunction");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Collector Policy resource
 *
 * @summary Creates or updates a Collector Policy resource
 * x-ms-original-file: specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-05-01/examples/CollectorPolicyCreate.json
 */
async function createACollectionPolicy() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const azureTrafficCollectorName = "atc";
  const collectorPolicyName = "cp1";
  const credential = new DefaultAzureCredential();
  const client = new AzureTrafficCollectorClient(credential, subscriptionId);
  const result = await client.collectorPolicies.beginCreateOrUpdateAndWait(
    resourceGroupName,
    azureTrafficCollectorName,
    collectorPolicyName
  );
  console.log(result);
}

createACollectionPolicy().catch(console.error);
