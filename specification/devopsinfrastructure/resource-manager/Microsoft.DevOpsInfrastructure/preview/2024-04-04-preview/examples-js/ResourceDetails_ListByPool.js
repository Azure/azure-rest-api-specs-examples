const { ManagedDevOpsInfrastructure } = require("@azure/arm-devopsinfrastructure");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List ResourceDetailsObject resources by Pool
 *
 * @summary List ResourceDetailsObject resources by Pool
 * x-ms-original-file: specification/devopsinfrastructure/resource-manager/Microsoft.DevOpsInfrastructure/preview/2024-04-04-preview/examples/ResourceDetails_ListByPool.json
 */
async function resourceDetailsListByPool() {
  const subscriptionId =
    process.env["DEVOPSINFRASTRUCTURE_SUBSCRIPTION_ID"] || "a2e95d27-c161-4b61-bda4-11512c14c2c2";
  const resourceGroupName =
    process.env["DEVOPSINFRASTRUCTURE_RESOURCE_GROUP"] || "my-resource-group";
  const poolName = "my-dev-ops-pool";
  const credential = new DefaultAzureCredential();
  const client = new ManagedDevOpsInfrastructure(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.resourceDetails.listByPool(resourceGroupName, poolName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
