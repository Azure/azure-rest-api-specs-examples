const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to List BuilderResource resources by subscription ID
 *
 * @summary List BuilderResource resources by subscription ID
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2025-02-02-preview/examples/Builders_ListBySubscription.json
 */
async function buildersListBySubscription0() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.builders.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}
