const { RelayAPI } = require("@azure/arm-relay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a service hybrid connection. This operation is idempotent.
 *
 * @summary Creates or updates a service hybrid connection. This operation is idempotent.
 * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/HybridConnection/RelayHybridConnectionCreate.json
 */
async function relayHybridConnectionCreate() {
  const subscriptionId =
    process.env["RELAY_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["RELAY_RESOURCE_GROUP"] || "resourcegroup";
  const namespaceName = "example-RelayNamespace-01";
  const hybridConnectionName = "example-Relay-Hybrid-01";
  const parameters = { requiresClientAuthorization: true };
  const credential = new DefaultAzureCredential();
  const client = new RelayAPI(credential, subscriptionId);
  const result = await client.hybridConnections.createOrUpdate(
    resourceGroupName,
    namespaceName,
    hybridConnectionName,
    parameters
  );
  console.log(result);
}
