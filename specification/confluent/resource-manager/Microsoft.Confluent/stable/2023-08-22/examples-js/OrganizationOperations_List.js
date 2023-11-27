const { ConfluentManagementClient } = require("@azure/arm-confluent");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all operations provided by Microsoft.Confluent.
 *
 * @summary List all operations provided by Microsoft.Confluent.
 * x-ms-original-file: specification/confluent/resource-manager/Microsoft.Confluent/stable/2023-08-22/examples/OrganizationOperations_List.json
 */
async function organizationOperationsList() {
  const credential = new DefaultAzureCredential();
  const client = new ConfluentManagementClient(credential);
  const resArray = new Array();
  for await (let item of client.organizationOperations.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}
