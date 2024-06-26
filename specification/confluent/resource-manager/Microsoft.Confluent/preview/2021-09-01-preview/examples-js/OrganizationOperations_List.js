const { ConfluentManagementClient } = require("@azure/arm-confluent");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all operations provided by Microsoft.Confluent.
 *
 * @summary List all operations provided by Microsoft.Confluent.
 * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/preview/2021-09-01-preview/examples/OrganizationOperations_List.json
 */
async function organizationOperationsList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new ConfluentManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.organizationOperations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

organizationOperationsList().catch(console.error);
