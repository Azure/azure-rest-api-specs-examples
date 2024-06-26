const { ConfluentManagementClient } = require("@azure/arm-confluent");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update Organization resource
 *
 * @summary Update Organization resource
 * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/preview/2021-09-01-preview/examples/Organization_Update.json
 */
async function confluentUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const organizationName = "myOrganization";
  const body = {
    tags: { client: "dev-client", env: "dev" },
  };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new ConfluentManagementClient(credential, subscriptionId);
  const result = await client.organization.update(resourceGroupName, organizationName, options);
  console.log(result);
}

confluentUpdate().catch(console.error);
