const { MicrosoftResourceHealth } = require("@azure/arm-resourcehealth");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets current availability status for a single resource
 *
 * @summary Gets current availability status for a single resource
 * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/preview/2023-10-01-preview/examples/AvailabilityStatus_GetByResource.json
 */
async function getCurrentHealthByResource() {
  const resourceUri = "resourceUri";
  const expand = "recommendedactions";
  const options = { expand };
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftResourceHealth(credential);
  const result = await client.availabilityStatuses.getByResource(resourceUri, options);
  console.log(result);
}
