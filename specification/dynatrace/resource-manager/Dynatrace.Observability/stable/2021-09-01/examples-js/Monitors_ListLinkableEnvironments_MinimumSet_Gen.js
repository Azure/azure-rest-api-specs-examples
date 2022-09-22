const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the Dynatrace environments that a user can link a azure resource to
 *
 * @summary Gets all the Dynatrace environments that a user can link a azure resource to
 * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/Monitors_ListLinkableEnvironments_MinimumSet_Gen.json
 */
async function monitorsListLinkableEnvironmentsMinimumSetGen() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "myMonitor";
  const request = {};
  const credential = new DefaultAzureCredential();
  const client = new DynatraceObservability(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.monitors.listLinkableEnvironments(
    resourceGroupName,
    monitorName,
    request
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

monitorsListLinkableEnvironmentsMinimumSetGen().catch(console.error);
