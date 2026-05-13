
import com.azure.resourcemanager.dataprotection.models.AlertsState;
import com.azure.resourcemanager.dataprotection.models.AzureMonitorAlertSettings;
import com.azure.resourcemanager.dataprotection.models.BackupVault;
import com.azure.resourcemanager.dataprotection.models.CrossRegionRestoreSettings;
import com.azure.resourcemanager.dataprotection.models.CrossRegionRestoreState;
import com.azure.resourcemanager.dataprotection.models.CrossSubscriptionRestoreSettings;
import com.azure.resourcemanager.dataprotection.models.CrossSubscriptionRestoreState;
import com.azure.resourcemanager.dataprotection.models.FeatureSettings;
import com.azure.resourcemanager.dataprotection.models.ImmutabilitySettings;
import com.azure.resourcemanager.dataprotection.models.ImmutabilityState;
import com.azure.resourcemanager.dataprotection.models.MonitoringSettings;
import com.azure.resourcemanager.dataprotection.models.SecuritySettings;
import com.azure.resourcemanager.dataprotection.models.SoftDeleteSettings;
import com.azure.resourcemanager.dataprotection.models.SoftDeleteState;
import com.azure.resourcemanager.dataprotection.models.StorageSetting;
import com.azure.resourcemanager.dataprotection.models.StorageSettingStoreTypes;
import com.azure.resourcemanager.dataprotection.models.StorageSettingTypes;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for BackupVaults CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/PutBackupVaultWithUndelete.json
     */
    /**
     * Sample code: Restore a soft-deleted backup vault.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        restoreASoftDeletedBackupVault(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupVaults().define("swaggerExample").withRegion("WestUS")
            .withExistingResourceGroup("SampleResourceGroup")
            .withProperties(new BackupVault()
                .withMonitoringSettings(new MonitoringSettings().withAzureMonitorAlertSettings(
                    new AzureMonitorAlertSettings().withAlertsForAllJobFailures(AlertsState.ENABLED)))
                .withSecuritySettings(new SecuritySettings()
                    .withSoftDeleteSettings(new SoftDeleteSettings().withState(SoftDeleteState.fromString("Enabled"))
                        .withRetentionDurationInDays(14.0D))
                    .withImmutabilitySettings(new ImmutabilitySettings().withState(ImmutabilityState.DISABLED)))
                .withStorageSettings(
                    Arrays.asList(new StorageSetting().withDatastoreType(StorageSettingStoreTypes.VAULT_STORE)
                        .withType(StorageSettingTypes.LOCALLY_REDUNDANT)))
                .withFeatureSettings(new FeatureSettings()
                    .withCrossSubscriptionRestoreSettings(
                        new CrossSubscriptionRestoreSettings().withState(CrossSubscriptionRestoreState.DISABLED))
                    .withCrossRegionRestoreSettings(
                        new CrossRegionRestoreSettings().withState(CrossRegionRestoreState.ENABLED))))
            .withTags(mapOf("key1", "fakeTokenPlaceholder"))
            .withXMsDeletedVaultId(
                "/subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.DataProtection/locations/WestUS/deletedVaults/swaggerExample")
            .create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
