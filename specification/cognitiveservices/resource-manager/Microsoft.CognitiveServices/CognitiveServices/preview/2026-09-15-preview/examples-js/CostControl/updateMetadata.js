const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates selected properties of a cost control.
 *
 * @summary updates selected properties of a cost control.
 * x-ms-original-file: 2026-09-15-preview/CostControl/updateMetadata.json
 */
async function updateMetadataWhilePreservingLegacyRules() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-1111-2222-3333-444444444444";
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.costControls.update(
    "foundry-resource-group",
    "foundry-account",
    "legacy-budget",
    { properties: { displayName: "Renamed legacy budget" } },
    { ifMatch: '"00000000-0000-0000-0000-000000000001"' },
  );
  console.log(result);
}
