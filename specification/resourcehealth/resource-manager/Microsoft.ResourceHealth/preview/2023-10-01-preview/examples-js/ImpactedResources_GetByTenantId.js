const { MicrosoftResourceHealth } = require("@azure/arm-resourcehealth");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specific impacted resource in the tenant by an event.
 *
 * @summary Gets the specific impacted resource in the tenant by an event.
 * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/ImpactedResources_GetByTenantId.json
 */
async function impactedResourcesGet() {
  const eventTrackingId = "BC_1-FXZ";
  const impactedResourceName = "abc-123-ghj-456";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftResourceHealth(credential);
  const result = await client.impactedResources.getByTenantId(
    eventTrackingId,
    impactedResourceName
  );
  console.log(result);
}
