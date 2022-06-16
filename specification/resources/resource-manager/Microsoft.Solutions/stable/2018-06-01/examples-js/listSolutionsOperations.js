const { ApplicationClient } = require("@azure/arm-managedapplications");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the available Microsoft.Solutions REST API operations.
 *
 * @summary Lists all of the available Microsoft.Solutions REST API operations.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/listSolutionsOperations.json
 */
async function listSolutionsOperations() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.listOperations()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listSolutionsOperations().catch(console.error);
