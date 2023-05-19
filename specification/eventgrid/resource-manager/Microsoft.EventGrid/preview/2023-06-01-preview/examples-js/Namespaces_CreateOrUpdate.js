const { EventGridManagementClient } = require("@azure/arm-eventgrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Asynchronously creates or updates a new namespace with the specified parameters.
 *
 * @summary Asynchronously creates or updates a new namespace with the specified parameters.
 * x-ms-original-file: specification/eventgrid/resource-manager/Microsoft.EventGrid/preview/2023-06-01-preview/examples/Namespaces_CreateOrUpdate.json
 */
async function namespacesCreateOrUpdate() {
  const subscriptionId =
    process.env["EVENTGRID_SUBSCRIPTION_ID"] || "8f6b6269-84f2-4d09-9e31-1127efcd1e40";
  const resourceGroupName = process.env["EVENTGRID_RESOURCE_GROUP"] || "examplerg";
  const namespaceName = "exampleNamespaceName1";
  const namespaceInfo = {
    location: "westus",
    tags: { tag1: "value11", tag2: "value22" },
    topicSpacesConfiguration: {
      routeTopicResourceId:
        "/subscriptions/8f6b6269-84f2-4d09-9e31-1127efcd1e40/resourceGroups/examplerg/providers/Microsoft.EventGrid/topics/exampleTopic1",
      state: "Enabled",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new EventGridManagementClient(credential, subscriptionId);
  const result = await client.namespaces.beginCreateOrUpdateAndWait(
    resourceGroupName,
    namespaceName,
    namespaceInfo
  );
  console.log(result);
}
