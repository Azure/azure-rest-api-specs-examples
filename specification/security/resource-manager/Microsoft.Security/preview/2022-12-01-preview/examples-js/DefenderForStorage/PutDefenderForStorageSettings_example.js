const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the Defender for Storage settings on a specified storage account.
 *
 * @summary Creates or updates the Defender for Storage settings on a specified storage account.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-12-01-preview/examples/DefenderForStorage/PutDefenderForStorageSettings_example.json
 */
async function createsOrUpdatesTheDefenderForStorageSettingsOnASpecifiedResource() {
  const resourceId =
    "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Storage/storageAccounts/samplestorageaccount";
  const settingName = "current";
  const defenderForStorageSetting = {
    capGBPerMonth: -1,
    isEnabledPropertiesMalwareScanningOnUploadIsEnabled: true,
    overrideSubscriptionLevelSettings: true,
    scanResultsEventGridTopicResourceId:
      "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.EventGrid/topics/sampletopic",
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential);
  const result = await client.defenderForStorage.create(
    resourceId,
    settingName,
    defenderForStorageSetting,
  );
  console.log(result);
}
