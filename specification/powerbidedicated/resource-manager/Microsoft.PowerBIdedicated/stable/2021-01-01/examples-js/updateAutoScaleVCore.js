const { PowerBIDedicated } = require("@azure/arm-powerbidedicated");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the current state of the specified auto scale v-core.
 *
 * @summary Updates the current state of the specified auto scale v-core.
 * x-ms-original-file: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/updateAutoScaleVCore.json
 */
async function updateAutoScaleVCoreParameters() {
  const subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
  const resourceGroupName = "TestRG";
  const vcoreName = "testvcore";
  const vCoreUpdateParameters = {
    capacityLimit: 20,
    sku: { name: "AutoScale", capacity: 0, tier: "AutoScale" },
    tags: { testKey: "testValue" },
  };
  const credential = new DefaultAzureCredential();
  const client = new PowerBIDedicated(credential, subscriptionId);
  const result = await client.autoScaleVCores.update(
    resourceGroupName,
    vcoreName,
    vCoreUpdateParameters
  );
  console.log(result);
}

updateAutoScaleVCoreParameters().catch(console.error);
