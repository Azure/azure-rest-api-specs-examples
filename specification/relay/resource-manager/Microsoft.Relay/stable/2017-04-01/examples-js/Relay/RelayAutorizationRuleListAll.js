const { RelayAPI } = require("@azure/arm-relay");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Authorization rules for a WCF relay.
 *
 * @summary Authorization rules for a WCF relay.
 * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/Relay/RelayAutorizationRuleListAll.json
 */
async function relayAutorizationRuleListAll() {
  const subscriptionId =
    process.env["RELAY_SUBSCRIPTION_ID"] || "ffffffff-ffff-ffff-ffff-ffffffffffff";
  const resourceGroupName = process.env["RELAY_RESOURCE_GROUP"] || "resourcegroup";
  const namespaceName = "example-RelayNamespace-01";
  const relayName = "example-Relay-Wcf-01";
  const credential = new DefaultAzureCredential();
  const client = new RelayAPI(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.wCFRelays.listAuthorizationRules(
    resourceGroupName,
    namespaceName,
    relayName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
