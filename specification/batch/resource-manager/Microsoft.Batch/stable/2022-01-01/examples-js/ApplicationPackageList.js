const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the application packages in the specified application.
 *
 * @summary Lists all of the application packages in the specified application.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/ApplicationPackageList.json
 */
async function applicationPackageList() {
  const subscriptionId = "subid";
  const resourceGroupName = "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const applicationName = "app1";
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.applicationPackageOperations.list(
    resourceGroupName,
    accountName,
    applicationName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

applicationPackageList().catch(console.error);
