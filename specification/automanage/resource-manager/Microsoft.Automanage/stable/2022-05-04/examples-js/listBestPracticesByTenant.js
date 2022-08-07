const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve a list of Automanage best practices
 *
 * @summary Retrieve a list of Automanage best practices
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listBestPracticesByTenant.json
 */
async function listAutomanageBestPractices() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new AutomanageClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.bestPractices.listByTenant()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAutomanageBestPractices().catch(console.error);
