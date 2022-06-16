const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an authorization rule for a namespace.
 *
 * @summary Creates or updates an authorization rule for a namespace.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/SBNameSpaceAuthorizationRuleCreate.json
 */
async function nameSpaceAuthorizationRuleCreate() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "ArunMonocle";
  const namespaceName = "sdk-Namespace-6914";
  const authorizationRuleName = "sdk-AuthRules-1788";
  const parameters = { rights: ["Listen", "Send"] };
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.namespaces.createOrUpdateAuthorizationRule(
    resourceGroupName,
    namespaceName,
    authorizationRuleName,
    parameters
  );
  console.log(result);
}

nameSpaceAuthorizationRuleCreate().catch(console.error);
