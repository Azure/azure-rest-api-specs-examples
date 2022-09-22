const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a DynatraceSingleSignOnResource
 *
 * @summary Create a DynatraceSingleSignOnResource
 * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/SingleSignOn_CreateOrUpdate_MinimumSet_Gen.json
 */
async function singleSignOnCreateOrUpdateMinimumSetGen() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "myMonitor";
  const configurationName = "default";
  const resource = {
    aadDomains: ["mpliftrdt20210811outlook.onmicrosoft.com"],
    singleSignOnUrl: "https://www.dynatrace.io",
  };
  const credential = new DefaultAzureCredential();
  const client = new DynatraceObservability(credential, subscriptionId);
  const result = await client.singleSignOn.beginCreateOrUpdateAndWait(
    resourceGroupName,
    monitorName,
    configurationName,
    resource
  );
  console.log(result);
}

singleSignOnCreateOrUpdateMinimumSetGen().catch(console.error);
