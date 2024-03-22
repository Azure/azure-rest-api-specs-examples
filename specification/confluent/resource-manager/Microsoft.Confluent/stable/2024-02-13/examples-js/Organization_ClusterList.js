const { ConfluentManagementClient } = require("@azure/arm-confluent");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists of all the clusters in a environment
 *
 * @summary Lists of all the clusters in a environment
 * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Organization_ClusterList.json
 */
async function organizationListClusters() {
  const subscriptionId =
    process.env["CONFLUENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONFLUENT_RESOURCE_GROUP"] || "myResourceGroup";
  const organizationName = "myOrganization";
  const environmentId = "env-12132";
  const pageSize = 10;
  const options = { pageSize };
  const credential = new DefaultAzureCredential();
  const client = new ConfluentManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.organization.listClusters(
    resourceGroupName,
    organizationName,
    environmentId,
    options,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
