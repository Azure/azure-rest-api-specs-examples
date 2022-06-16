const { PowerBIDedicated } = require("@azure/arm-powerbidedicated");
const { DefaultAzureCredential } = require("@azure/identity");

async function getDetailsOfAnAutoScaleVCore() {
  const subscriptionId = "613192d7-503f-477a-9cfe-4efc3ee2bd60";
  const resourceGroupName = "TestRG";
  const vcoreName = "testvcore";
  const credential = new DefaultAzureCredential();
  const client = new PowerBIDedicated(credential, subscriptionId);
  const result = await client.autoScaleVCores.get(resourceGroupName, vcoreName);
  console.log(result);
}

getDetailsOfAnAutoScaleVCore().catch(console.error);
