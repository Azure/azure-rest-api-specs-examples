const { ConfluentManagementClient } = require("@azure/arm-confluent");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete Organization resource
 *
 * @summary Delete Organization resource
 * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/preview/2021-09-01-preview/examples/Organization_Delete.json
 */
async function confluentDelete() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const organizationName = "myOrganization";
  const credential = new DefaultAzureCredential();
  const client = new ConfluentManagementClient(credential, subscriptionId);
  const result = await client.organization.beginDeleteAndWait(resourceGroupName, organizationName);
  console.log(result);
}

confluentDelete().catch(console.error);
