const { ConfluentManagementClient } = require("@azure/arm-confluent");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get schema registry cluster by Id
 *
 * @summary Get schema registry cluster by Id
 * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Organization_GetSchemaRegistryClusterById.json
 */
async function organizationGetSchemaRegistryClusterById() {
  const subscriptionId =
    process.env["CONFLUENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONFLUENT_RESOURCE_GROUP"] || "myResourceGroup";
  const organizationName = "myOrganization";
  const environmentId = "env-stgcczjp2j3";
  const clusterId = "lsrc-stgczkq22z";
  const credential = new DefaultAzureCredential();
  const client = new ConfluentManagementClient(credential, subscriptionId);
  const result = await client.organization.getSchemaRegistryClusterById(
    resourceGroupName,
    organizationName,
    environmentId,
    clusterId,
  );
  console.log(result);
}
