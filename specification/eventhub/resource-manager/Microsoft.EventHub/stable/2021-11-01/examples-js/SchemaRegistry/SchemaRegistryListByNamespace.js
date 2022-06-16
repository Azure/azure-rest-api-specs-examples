const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the Schema Groups in a Namespace.
 *
 * @summary Gets all the Schema Groups in a Namespace.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/SchemaRegistry/SchemaRegistryListByNamespace.json
 */
async function schemaRegistryListAll() {
  const subscriptionId = "e8baea74-64ce-459b-bee3-5aa4c47b3ae3";
  const resourceGroupName = "alitest";
  const namespaceName = "ali-ua-test-eh-system-1";
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.schemaRegistry.listByNamespace(resourceGroupName, namespaceName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

schemaRegistryListAll().catch(console.error);
