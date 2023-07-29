const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the applications in the specified account.
 *
 * @summary Lists all of the applications in the specified account.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2023-05-01/examples/ApplicationList.json
 */
async function applicationList() {
  const subscriptionId = process.env["BATCH_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["BATCH_RESOURCE_GROUP"] || "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.applicationOperations.list(resourceGroupName, accountName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
