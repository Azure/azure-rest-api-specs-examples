const { LoadTestClient } = require("@azure/arm-loadtesting");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the available quota per region per subscription.
 *
 * @summary Lists all the available quota per region per subscription.
 * x-ms-original-file: specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/Quotas_List.json
 */
async function quotasList() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const location = "westus";
  const credential = new DefaultAzureCredential();
  const client = new LoadTestClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.quotas.list(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}

quotasList().catch(console.error);
