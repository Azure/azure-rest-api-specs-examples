const { ManagedDevOpsInfrastructure } = require("@azure/arm-devopsinfrastructure");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List Pool resources by resource group
 *
 * @summary List Pool resources by resource group
 * x-ms-original-file: specification/devopsinfrastructure/resource-manager/Microsoft.DevOpsInfrastructure/preview/2024-04-04-preview/examples/ListPoolsBySubscriptionAndResourceGroup.json
 */
async function poolsListByResourceGroup() {
  const subscriptionId =
    process.env["DEVOPSINFRASTRUCTURE_SUBSCRIPTION_ID"] || "a2e95d27-c161-4b61-bda4-11512c14c2c2";
  const resourceGroupName = process.env["DEVOPSINFRASTRUCTURE_RESOURCE_GROUP"] || "rg";
  const credential = new DefaultAzureCredential();
  const client = new ManagedDevOpsInfrastructure(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.pools.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
