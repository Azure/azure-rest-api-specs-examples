const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List all DynatraceSingleSignOnResource by monitorName
 *
 * @summary List all DynatraceSingleSignOnResource by monitorName
 * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/SingleSignOn_List_MaximumSet_Gen.json
 */
async function singleSignOnListMaximumSetGen() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "myMonitor";
  const credential = new DefaultAzureCredential();
  const client = new DynatraceObservability(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.singleSignOn.list(resourceGroupName, monitorName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

singleSignOnListMaximumSetGen().catch(console.error);
