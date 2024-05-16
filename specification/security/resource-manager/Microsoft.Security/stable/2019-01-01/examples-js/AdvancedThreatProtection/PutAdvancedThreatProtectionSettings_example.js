const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the Advanced Threat Protection settings on a specified resource.
 *
 * @summary Creates or updates the Advanced Threat Protection settings on a specified resource.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-01-01/examples/AdvancedThreatProtection/PutAdvancedThreatProtectionSettings_example.json
 */
async function createsOrUpdatesTheAdvancedThreatProtectionSettingsOnASpecifiedResource() {
  const resourceId =
    "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Storage/storageAccounts/samplestorageaccount";
  const advancedThreatProtectionSetting = {
    name: "current",
    type: "Microsoft.Security/advancedThreatProtectionSettings",
    id: "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Storage/storageAccounts/samplestorageaccount/providers/Microsoft.Security/advancedThreatProtectionSettings/current",
    isEnabled: true,
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const result = await client.advancedThreatProtection.create(
    resourceId,
    advancedThreatProtectionSetting,
  );
  console.log(result);
}
