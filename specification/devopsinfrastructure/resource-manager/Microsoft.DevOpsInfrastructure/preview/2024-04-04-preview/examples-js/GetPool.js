const { ManagedDevOpsInfrastructure } = require("@azure/arm-devopsinfrastructure");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a Pool
 *
 * @summary Get a Pool
 * x-ms-original-file: specification/devopsinfrastructure/resource-manager/Microsoft.DevOpsInfrastructure/preview/2024-04-04-preview/examples/GetPool.json
 */
async function poolsGet() {
  const subscriptionId =
    process.env["DEVOPSINFRASTRUCTURE_SUBSCRIPTION_ID"] || "a2e95d27-c161-4b61-bda4-11512c14c2c2";
  const resourceGroupName = process.env["DEVOPSINFRASTRUCTURE_RESOURCE_GROUP"] || "rg";
  const poolName = "pool";
  const credential = new DefaultAzureCredential();
  const client = new ManagedDevOpsInfrastructure(credential, subscriptionId);
  const result = await client.pools.get(resourceGroupName, poolName);
  console.log(result);
}
