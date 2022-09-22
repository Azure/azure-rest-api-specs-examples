const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a DynatraceSingleSignOnResource
 *
 * @summary Get a DynatraceSingleSignOnResource
 * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/SingleSignOn_Get_MaximumSet_Gen.json
 */
async function singleSignOnGetMaximumSetGen() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "myMonitor";
  const configurationName = "default";
  const credential = new DefaultAzureCredential();
  const client = new DynatraceObservability(credential, subscriptionId);
  const result = await client.singleSignOn.get(resourceGroupName, monitorName, configurationName);
  console.log(result);
}

singleSignOnGetMaximumSetGen().catch(console.error);
