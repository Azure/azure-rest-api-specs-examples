const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve a list of Automanage best practices versions
 *
 * @summary Retrieve a list of Automanage best practices versions
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listBestPracticesVersionsByTenant.json
 */
async function listAutomanageBestPracticesVersions() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const bestPracticeName = "azureBestPracticesProduction";
  const credential = new DefaultAzureCredential();
  const client = new AutomanageClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.bestPracticesVersions.listByTenant(bestPracticeName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAutomanageBestPracticesVersions().catch(console.error);
