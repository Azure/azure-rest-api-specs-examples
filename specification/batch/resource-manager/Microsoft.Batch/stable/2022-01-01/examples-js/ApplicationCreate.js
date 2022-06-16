const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Adds an application to the specified Batch account.
 *
 * @summary Adds an application to the specified Batch account.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/ApplicationCreate.json
 */
async function applicationCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "default-azurebatch-japaneast";
  const accountName = "sampleacct";
  const applicationName = "app1";
  const parameters = {
    allowUpdates: false,
    displayName: "myAppName",
  };
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.applicationOperations.create(
    resourceGroupName,
    accountName,
    applicationName,
    options
  );
  console.log(result);
}

applicationCreate().catch(console.error);
