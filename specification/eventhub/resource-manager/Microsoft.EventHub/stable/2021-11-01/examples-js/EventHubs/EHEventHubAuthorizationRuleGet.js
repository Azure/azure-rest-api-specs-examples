const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an AuthorizationRule for an Event Hub by rule name.
 *
 * @summary Gets an AuthorizationRule for an Event Hub by rule name.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/EventHubs/EHEventHubAuthorizationRuleGet.json
 */
async function eventHubAuthorizationRuleGet() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "ArunMonocle";
  const namespaceName = "sdk-Namespace-960";
  const eventHubName = "sdk-EventHub-532";
  const authorizationRuleName = "sdk-Authrules-2513";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.eventHubs.getAuthorizationRule(
    resourceGroupName,
    namespaceName,
    eventHubName,
    authorizationRuleName
  );
  console.log(result);
}

eventHubAuthorizationRuleGet().catch(console.error);
