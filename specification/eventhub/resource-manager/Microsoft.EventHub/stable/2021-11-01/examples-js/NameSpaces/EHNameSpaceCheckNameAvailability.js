const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check the give Namespace name availability.
 *
 * @summary Check the give Namespace name availability.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/EHNameSpaceCheckNameAvailability.json
 */
async function namespacesCheckNameAvailability() {
  const subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
  const parameters = {
    name: "sdk-Namespace-8458",
  };
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.namespaces.checkNameAvailability(parameters);
  console.log(result);
}

namespacesCheckNameAvailability().catch(console.error);
