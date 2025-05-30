const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates an ApplicationGroup for a Namespace.
 *
 * @summary Creates or updates an ApplicationGroup for a Namespace.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2024-05-01-preview/examples/ApplicationGroup/ApplicationGroupCreate.json
 */
async function applicationGroupCreate() {
  const subscriptionId =
    process.env["EVENTHUB_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["EVENTHUB_RESOURCE_GROUP"] || "contosotest";
  const namespaceName = "contoso-ua-test-eh-system-1";
  const applicationGroupName = "appGroup1";
  const parameters = {
    clientAppGroupIdentifier: "SASKeyName=KeyName",
    isEnabled: true,
    policies: [
      {
        name: "ThrottlingPolicy1",
        type: "ThrottlingPolicy",
        metricId: "IncomingMessages",
        rateLimitThreshold: 7912,
      },
      {
        name: "ThrottlingPolicy2",
        type: "ThrottlingPolicy",
        metricId: "IncomingBytes",
        rateLimitThreshold: 3951729,
      },
      {
        name: "ThrottlingPolicy3",
        type: "ThrottlingPolicy",
        metricId: "OutgoingBytes",
        rateLimitThreshold: 245175,
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.applicationGroupOperations.createOrUpdateApplicationGroup(
    resourceGroupName,
    namespaceName,
    applicationGroupName,
    parameters,
  );
  console.log(result);
}
