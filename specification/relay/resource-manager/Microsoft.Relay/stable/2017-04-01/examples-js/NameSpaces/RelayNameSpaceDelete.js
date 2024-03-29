const { RelayAPI } = require("@azure/arm-relay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an existing namespace. This operation also removes all associated resources under the namespace.
 *
 * @summary Deletes an existing namespace. This operation also removes all associated resources under the namespace.
 * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/NameSpaces/RelayNameSpaceDelete.json
 */
async function relayNameSpaceDelete() {
  const subscriptionId =
    process.env["RELAY_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["RELAY_RESOURCE_GROUP"] || "resourcegroup";
  const namespaceName = "example-RelayNamespace-01";
  const credential = new DefaultAzureCredential();
  const client = new RelayAPI(credential, subscriptionId);
  const result = await client.namespaces.beginDeleteAndWait(resourceGroupName, namespaceName);
  console.log(result);
}
