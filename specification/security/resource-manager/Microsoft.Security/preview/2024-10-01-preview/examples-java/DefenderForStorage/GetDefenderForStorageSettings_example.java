
import com.azure.resourcemanager.security.models.SettingName;

/**
 * Samples for DefenderForStorage Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2024-10-01-preview/examples/DefenderForStorage
     * /GetDefenderForStorageSettings_example.json
     */
    /**
     * Sample code: Gets the Defender for Storage settings for the specified resource.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getsTheDefenderForStorageSettingsForTheSpecifiedResource(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.defenderForStorages().getWithResponse(
            "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Storage/storageAccounts/samplestorageaccount",
            SettingName.CURRENT, com.azure.core.util.Context.NONE);
    }
}
