const { ConfluentManagementClient } = require("@azure/arm-confluent");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes API key of a kafka or schema registry cluster
 *
 * @summary Deletes API key of a kafka or schema registry cluster
 * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Organization_DeleteClusterAPIKey.json
 */
async function organizationDeleteClusterApiKey() {
  const subscriptionId =
    process.env["CONFLUENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONFLUENT_RESOURCE_GROUP"] || "myResourceGroup";
  const organizationName = "myOrganization";
  const apiKeyId = "ZFZ6SZZZWGYBEIFB";
  const credential = new DefaultAzureCredential();
  const client = new ConfluentManagementClient(credential, subscriptionId);
  const result = await client.organization.deleteClusterAPIKey(
    resourceGroupName,
    organizationName,
    apiKeyId,
  );
  console.log(result);
}
