const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a DynatraceSingleSignOnResource
 *
 * @summary Create a DynatraceSingleSignOnResource
 * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/preview/2021-09-01-preview/examples/SingleSignOn_CreateOrUpdate_MaximumSet_Gen.json
 */
async function singleSignOnCreateOrUpdateMaximumSetGen() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "myMonitor";
  const configurationName = "default";
  const resource = {
    aadDomains: ["mpliftrdt20210811outlook.onmicrosoft.com"],
    enterpriseAppId: "00000000-0000-0000-0000-000000000000",
    provisioningState: "Accepted",
    singleSignOnState: "Enable",
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

singleSignOnCreateOrUpdateMaximumSetGen().catch(console.error);
