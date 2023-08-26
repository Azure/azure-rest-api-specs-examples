const { DynatraceObservability } = require("@azure/arm-dynatrace");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get Marketplace SaaS resource details of a tenant under a specific subscription
 *
 * @summary Get Marketplace SaaS resource details of a tenant under a specific subscription
 * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2023-04-27/examples/Monitors_GetMarketplaceSaaSResourceDetails_MaximumSet_Gen.json
 */
async function monitorsGetMarketplaceSaaSResourceDetailsMaximumSetGen() {
  const subscriptionId = process.env["DYNATRACE_SUBSCRIPTION_ID"] || "nqmcgifgaqlf";
  const request = {
    tenantId: "urnmattojzhktcfw",
  };
  const credential = new DefaultAzureCredential();
  const client = new DynatraceObservability(credential, subscriptionId);
  const result = await client.monitors.getMarketplaceSaaSResourceDetails(request);
  console.log(result);
}
