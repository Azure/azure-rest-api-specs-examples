const { ConfluentManagementClient } = require("@azure/arm-confluent");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to cloud provider regions available for creating Schema Registry clusters.
 *
 * @summary cloud provider regions available for creating Schema Registry clusters.
 * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2024-02-13/examples/Organization_ListRegions.json
 */
async function organizationListRegions() {
  const subscriptionId =
    process.env["CONFLUENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONFLUENT_RESOURCE_GROUP"] || "myResourceGroup";
  const organizationName = "myOrganization";
  const body = {
    searchFilters: {
      cloud: "azure",
      packages: "ADVANCED,ESSENTIALS",
      region: "eastus",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ConfluentManagementClient(credential, subscriptionId);
  const result = await client.organization.listRegions(resourceGroupName, organizationName, body);
  console.log(result);
}
