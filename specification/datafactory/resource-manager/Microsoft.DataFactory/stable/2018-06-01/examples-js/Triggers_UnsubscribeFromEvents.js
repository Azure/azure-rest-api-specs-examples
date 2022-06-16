const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Unsubscribe event trigger from events.
 *
 * @summary Unsubscribe event trigger from events.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_UnsubscribeFromEvents.json
 */
async function triggersUnsubscribeFromEvents() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const triggerName = "exampleTrigger";
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.triggers.beginUnsubscribeFromEventsAndWait(
    resourceGroupName,
    factoryName,
    triggerName
  );
  console.log(result);
}

triggersUnsubscribeFromEvents().catch(console.error);
