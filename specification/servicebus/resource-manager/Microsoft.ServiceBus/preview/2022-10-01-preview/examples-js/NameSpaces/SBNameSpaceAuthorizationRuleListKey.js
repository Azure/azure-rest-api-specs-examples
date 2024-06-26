const { ServiceBusManagementClient } = require("@azure/arm-servicebus");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the primary and secondary connection strings for the namespace.
 *
 * @summary Gets the primary and secondary connection strings for the namespace.
 * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/preview/2022-10-01-preview/examples/NameSpaces/SBNameSpaceAuthorizationRuleListKey.json
 */
async function nameSpaceAuthorizationRuleListKey() {
  const subscriptionId =
    process.env["SERVICEBUS_SUBSCRIPTION_ID"] || "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = process.env["SERVICEBUS_RESOURCE_GROUP"] || "ArunMonocle";
  const namespaceName = "sdk-namespace-6914";
  const authorizationRuleName = "sdk-AuthRules-1788";
  const credential = new DefaultAzureCredential();
  const client = new ServiceBusManagementClient(credential, subscriptionId);
  const result = await client.namespaces.listKeys(
    resourceGroupName,
    namespaceName,
    authorizationRuleName
  );
  console.log(result);
}
