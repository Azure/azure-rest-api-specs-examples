const { MicrosoftResourceHealth } = require("@azure/arm-resourcehealth");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists impacted resources in the tenant by an event (Security Advisory).
 *
 * @summary Lists impacted resources in the tenant by an event (Security Advisory).
 * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2022-10-01/examples/SecurityAdvisoryImpactedResources_ListByTenantId_ListByEventId.json
 */
async function listSecurityAdvisoryImpactedResourcesByTenantId() {
  const subscriptionId =
    process.env["RESOURCEHEALTH_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const eventTrackingId = "BC_1-FXZ";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftResourceHealth(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.securityAdvisoryImpactedResources.listByTenantIdAndEventId(
    eventTrackingId
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
