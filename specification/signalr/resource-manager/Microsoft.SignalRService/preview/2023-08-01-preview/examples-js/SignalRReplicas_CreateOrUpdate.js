const { SignalRManagementClient } = require("@azure/arm-signalr");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a replica.
 *
 * @summary Create or update a replica.
 * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/preview/2023-08-01-preview/examples/SignalRReplicas_CreateOrUpdate.json
 */
async function signalRReplicasCreateOrUpdate() {
  const subscriptionId =
    process.env["SIGNALR_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["SIGNALR_RESOURCE_GROUP"] || "myResourceGroup";
  const resourceName = "mySignalRService";
  const replicaName = "mySignalRService-eastus";
  const parameters = {
    location: "eastus",
    resourceStopped: "false",
    sku: { name: "Premium_P1", capacity: 1, tier: "Premium" },
    tags: { key1: "value1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new SignalRManagementClient(credential, subscriptionId);
  const result = await client.signalRReplicas.beginCreateOrUpdateAndWait(
    resourceGroupName,
    resourceName,
    replicaName,
    parameters
  );
  console.log(result);
}
