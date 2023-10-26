const { WebPubSubManagementClient } = require("@azure/arm-webpubsub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a replica.
 *
 * @summary Create or update a replica.
 * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/WebPubSubReplicas_CreateOrUpdate.json
 */
async function webPubSubReplicasCreateOrUpdate() {
  const subscriptionId =
    process.env["WEB-PUBSUB_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["WEB-PUBSUB_RESOURCE_GROUP"] || "myResourceGroup";
  const resourceName = "myWebPubSubService";
  const replicaName = "myWebPubSubService-eastus";
  const parameters = {
    location: "eastus",
    resourceStopped: "false",
    sku: { name: "Premium_P1", capacity: 1, tier: "Premium" },
    tags: { key1: "value1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new WebPubSubManagementClient(credential, subscriptionId);
  const result = await client.webPubSubReplicas.beginCreateOrUpdateAndWait(
    resourceGroupName,
    resourceName,
    replicaName,
    parameters
  );
  console.log(result);
}
