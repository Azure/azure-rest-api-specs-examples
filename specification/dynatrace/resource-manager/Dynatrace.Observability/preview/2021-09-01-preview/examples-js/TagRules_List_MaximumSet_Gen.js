const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all TagRule by monitorName
 *
 * @summary List all TagRule by monitorName
 * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/TagRules_List_MaximumSet_Gen.json
 */
async function tagRulesListMaximumSetGen() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "myMonitor";
  const credential = new DefaultAzureCredential();
  const client = new DynatraceObservability(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.tagRules.list(resourceGroupName, monitorName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

tagRulesListMaximumSetGen().catch(console.error);
