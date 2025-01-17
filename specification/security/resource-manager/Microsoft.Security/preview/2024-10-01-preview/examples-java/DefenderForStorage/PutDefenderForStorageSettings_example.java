
import com.azure.resourcemanager.security.models.DefenderForStorageSettingProperties;
import com.azure.resourcemanager.security.models.MalwareScanningProperties;
import com.azure.resourcemanager.security.models.OnUploadProperties;
import com.azure.resourcemanager.security.models.SensitiveDataDiscoveryProperties;
import com.azure.resourcemanager.security.models.SettingName;

/**
 * Samples for DefenderForStorage Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2024-10-01-preview/examples/DefenderForStorage
     * /PutDefenderForStorageSettings_example.json
     */
    /**
     * Sample code: Creates or updates the Defender for Storage settings on a specified resource.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void createsOrUpdatesTheDefenderForStorageSettingsOnASpecifiedResource(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.defenderForStorages().define(SettingName.CURRENT).withExistingResourceId(
            "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Storage/storageAccounts/samplestorageaccount")
            .withProperties(new DefenderForStorageSettingProperties().withIsEnabled(true)
                .withMalwareScanning(new MalwareScanningProperties()
                    .withOnUpload(new OnUploadProperties().withIsEnabled(true).withCapGBPerMonth(-1))
                    .withScanResultsEventGridTopicResourceId(
                        "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.EventGrid/topics/sampletopic"))
                .withSensitiveDataDiscovery(new SensitiveDataDiscoveryProperties().withIsEnabled(true))
                .withOverrideSubscriptionLevelSettings(true))
            .create();
    }
}
