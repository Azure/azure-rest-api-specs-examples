const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to List all the namespace topics under a namespace.
 *
 * @summary List all the namespace topics under a namespace.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/stable/2025-02-15/examples/NamespaceTopics_ListByNamespace.json
 */
async function namespaceTopicsListByNamespace() {
  const subscriptionId =
    process.env["EVENTGRID_SUBSCRIPTION_ID"] || "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
  const resourceGroupName = process.env["EVENTGRID_RESOURCE_GROUP"] || "examplerg";
  const namespaceName = "examplenamespace2";
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.namespaceTopics.listByNamespace(
    resourceGroupName,
    namespaceName,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
