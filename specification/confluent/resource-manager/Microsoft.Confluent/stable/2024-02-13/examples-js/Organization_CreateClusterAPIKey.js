const { ConfluentManagementClient } = require("@azure/arm-confluent");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates API key for a schema registry Cluster ID or Kafka Cluster ID under a environment
 *
 * @summary Creates API key for a schema registry Cluster ID or Kafka Cluster ID under a environment
 * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Organization_CreateClusterAPIKey.json
 */
async function organizationCreateApiKey() {
  const subscriptionId =
    process.env["CONFLUENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONFLUENT_RESOURCE_GROUP"] || "myResourceGroup";
  const organizationName = "myOrganization";
  const environmentId = "env-12132";
  const clusterId = "clusterId-123";
  const body = {
    name: "CI kafka access key",
    description: "This API key provides kafka access to cluster x",
  };
  const credential = new DefaultAzureCredential();
  const client = new ConfluentManagementClient(credential, subscriptionId);
  const result = await client.organization.createAPIKey(
    resourceGroupName,
    organizationName,
    environmentId,
    clusterId,
    body,
  );
  console.log(result);
}
