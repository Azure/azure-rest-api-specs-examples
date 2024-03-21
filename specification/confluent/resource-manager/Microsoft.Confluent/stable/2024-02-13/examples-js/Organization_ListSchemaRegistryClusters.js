const { ConfluentManagementClient } = require("@azure/arm-confluent");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get schema registry clusters
 *
 * @summary Get schema registry clusters
 * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Organization_ListSchemaRegistryClusters.json
 */
async function organizationListSchemaRegistryClusters() {
  const subscriptionId =
    process.env["CONFLUENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONFLUENT_RESOURCE_GROUP"] || "myResourceGroup";
  const organizationName = "myOrganization";
  const environmentId = "env-stgcczjp2j3";
  const credential = new DefaultAzureCredential();
  const client = new ConfluentManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.organization.listSchemaRegistryClusters(
    resourceGroupName,
    organizationName,
    environmentId,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
