const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Subscribe event trigger to events.
 *
 * @summary Subscribe event trigger to events.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_SubscribeToEvents.json
 */
async function triggersSubscribeToEvents() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const triggerName = "exampleTrigger";
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.triggers.beginSubscribeToEventsAndWait(
    resourceGroupName,
    factoryName,
    triggerName
  );
  console.log(result);
}

triggersSubscribeToEvents().catch(console.error);
