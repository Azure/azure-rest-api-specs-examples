
import com.azure.resourcemanager.dataprotection.models.AzureBackupRecoveryPointBasedRestoreRequest;
import com.azure.resourcemanager.dataprotection.models.Datasource;
import com.azure.resourcemanager.dataprotection.models.DatasourceSet;
import com.azure.resourcemanager.dataprotection.models.IdentityDetails;
import com.azure.resourcemanager.dataprotection.models.RecoveryOption;
import com.azure.resourcemanager.dataprotection.models.RestoreTargetInfo;
import com.azure.resourcemanager.dataprotection.models.SecretStoreBasedAuthCredentials;
import com.azure.resourcemanager.dataprotection.models.SecretStoreResource;
import com.azure.resourcemanager.dataprotection.models.SecretStoreType;
import com.azure.resourcemanager.dataprotection.models.SourceDataStoreType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for BackupInstances TriggerRestore.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/BackupInstanceOperations/TriggerRestore.json
     */
    /**
     * Sample code: Trigger Restore.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void triggerRestore(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().triggerRestore("000pikumar", "PratikPrivatePreviewVault1", "testInstance1",
            new AzureBackupRecoveryPointBasedRestoreRequest().withRestoreTargetInfo(new RestoreTargetInfo()
                .withRecoveryOption(RecoveryOption.FAIL_IF_EXISTS).withRestoreLocation("southeastasia")
                .withDatasourceInfo(new Datasource().withDatasourceType("Microsoft.DBforPostgreSQL/servers/databases")
                    .withObjectType("Datasource")
                    .withResourceId(
                        "/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/targetdb")
                    .withResourceLocation("").withResourceName("targetdb")
                    .withResourceType("Microsoft.DBforPostgreSQL/servers/databases").withResourceUri(""))
                .withDatasourceSetInfo(new DatasourceSet()
                    .withDatasourceType("Microsoft.DBforPostgreSQL/servers/databases").withObjectType("DatasourceSet")
                    .withResourceId(
                        "/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest")
                    .withResourceLocation("").withResourceName("viveksipgtest")
                    .withResourceType("Microsoft.DBforPostgreSQL/servers").withResourceUri(""))
                .withDatasourceAuthCredentials(new SecretStoreBasedAuthCredentials().withSecretStoreResource(
                    new SecretStoreResource().withUri("https://samplevault.vault.azure.net/secrets/credentials")
                        .withSecretStoreType(SecretStoreType.AZURE_KEY_VAULT))))
                .withSourceDataStoreType(SourceDataStoreType.VAULT_STORE)
                .withSourceResourceId(
                    "/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb")
                .withIdentityDetails(
                    new IdentityDetails().withUseSystemAssignedIdentity(false).withUserAssignedIdentityArmUrl(
                        "/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/rg-name/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testUami"))
                .withRecoveryPointId("hardcodedRP"),
            com.azure.core.util.Context.NONE);
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
