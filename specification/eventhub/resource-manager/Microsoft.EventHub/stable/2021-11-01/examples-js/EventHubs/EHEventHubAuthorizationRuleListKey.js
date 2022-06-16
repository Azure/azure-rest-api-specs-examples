const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the ACS and SAS connection strings for the Event Hub.
 *
 * @summary Gets the ACS and SAS connection strings for the Event Hub.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/EventHubs/EHEventHubAuthorizationRuleListKey.json
 */
async function eventHubAuthorizationRuleListKey() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const resourceGroupName = "ArunMonocle";
  const namespaceName = "sdk-namespace-960";
  const eventHubName = "sdk-EventHub-532";
  const authorizationRuleName = "sdk-Authrules-2513";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.eventHubs.listKeys(
    resourceGroupName,
    namespaceName,
    eventHubName,
    authorizationRuleName
  );
  console.log(result);
}

eventHubAuthorizationRuleListKey().catch(console.error);
