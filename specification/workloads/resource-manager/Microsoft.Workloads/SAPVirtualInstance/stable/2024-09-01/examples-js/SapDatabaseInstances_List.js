const { WorkloadsClient } = require("@azure/arm-workloadssapvirtualinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to lists the Database resources associated with a Virtual Instance for SAP solutions resource.
 *
 * @summary lists the Database resources associated with a Virtual Instance for SAP solutions resource.
 * x-ms-original-file: 2024-09-01/SapDatabaseInstances_List.json
 */
async function sapDatabaseInstancesListBySAPVirtualInstance() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "6d875e77-e412-4d7d-9af4-8895278b4443";
  const client = new WorkloadsClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.sapDatabaseInstances.list("test-rg", "X00")) {
    resArray.push(item);
  }

  console.log(resArray);
}
