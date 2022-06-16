const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a trigger's event subscription status.
 *
 * @summary Get a trigger's event subscription status.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/Triggers_GetEventSubscriptionStatus.json
 */
async function triggersGetEventSubscriptionStatus() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const triggerName = "exampleTrigger";
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.triggers.getEventSubscriptionStatus(
    resourceGroupName,
    factoryName,
    triggerName
  );
  console.log(result);
}

triggersGetEventSubscriptionStatus().catch(console.error);
