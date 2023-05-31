const { QumuloStorage } = require("@azure/arm-qumulo");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List FileSystemResource resources by subscription ID
 *
 * @summary List FileSystemResource resources by subscription ID
 * x-ms-original-file: specification/liftrqumulo/resource-manager/Qumulo.Storage/stable/2022-10-12/examples/FileSystems_ListBySubscription_MinimumSet_Gen.json
 */
async function fileSystemsListBySubscriptionMinimumSetGen() {
  const subscriptionId = process.env["LIFTRQUMULO_SUBSCRIPTION_ID"] || "aaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new QumuloStorage(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.fileSystems.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
