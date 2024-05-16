const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the Advanced Threat Protection settings for the specified resource.
 *
 * @summary Gets the Advanced Threat Protection settings for the specified resource.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-01-01/examples/AdvancedThreatProtection/GetAdvancedThreatProtectionSettings_example.json
 */
async function getsTheAdvancedThreatProtectionSettingsForTheSpecifiedResource() {
  const resourceId =
    "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Storage/storageAccounts/samplestorageaccount";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const result = await client.advancedThreatProtection.get(resourceId);
  console.log(result);
}
