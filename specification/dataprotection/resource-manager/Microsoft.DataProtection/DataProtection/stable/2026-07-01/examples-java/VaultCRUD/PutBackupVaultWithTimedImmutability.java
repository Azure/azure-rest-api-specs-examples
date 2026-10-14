
import com.azure.resourcemanager.dataprotection.models.BackupVault;
import com.azure.resourcemanager.dataprotection.models.DppIdentityDetails;
import com.azure.resourcemanager.dataprotection.models.ImmutabilityConfiguration;
import com.azure.resourcemanager.dataprotection.models.ImmutabilitySettings;
import com.azure.resourcemanager.dataprotection.models.ImmutabilityState;
import com.azure.resourcemanager.dataprotection.models.ImmutabilityType;
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
     * x-ms-original-file: 2026-07-01/VaultCRUD/PutBackupVaultWithTimedImmutability.json
     */
    /**
     * Sample code: Create BackupVault With Timed Immutability.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        createBackupVaultWithTimedImmutability(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupVaults().define("swaggerExample").withRegion("WestUS")
            .withExistingResourceGroup("SampleResourceGroup")
            .withProperties(new BackupVault()
                .withSecuritySettings(new SecuritySettings()
                    .withSoftDeleteSettings(
                        new SoftDeleteSettings().withState(SoftDeleteState.ON).withRetentionDurationInDays(14.0D))
                    .withImmutabilitySettings(new ImmutabilitySettings().withState(ImmutabilityState.UNLOCKED)
                        .withConfiguration(new ImmutabilityConfiguration().withType(ImmutabilityType.TIME_BASED)
                            .withDurationInDays(30))))
                .withStorageSettings(
                    Arrays.asList(new StorageSetting().withDatastoreType(StorageSettingStoreTypes.VAULT_STORE)
                        .withType(StorageSettingTypes.LOCALLY_REDUNDANT))))
            .withTags(mapOf("key1", "fakeTokenPlaceholder")).withIdentity(new DppIdentityDetails().withType("None"))
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
