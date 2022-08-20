const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Partially updates a machine pool
 *
 * @summary Partially updates a machine pool
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/Pools_Patch.json
 */
async function poolsUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const projectName = "{projectName}";
  const poolName = "{poolName}";
  const body = { devBoxDefinitionName: "WebDevBox2" };
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.pools.beginUpdateAndWait(
    resourceGroupName,
    projectName,
    poolName,
    body
  );
  console.log(result);
}

poolsUpdate().catch(console.error);
