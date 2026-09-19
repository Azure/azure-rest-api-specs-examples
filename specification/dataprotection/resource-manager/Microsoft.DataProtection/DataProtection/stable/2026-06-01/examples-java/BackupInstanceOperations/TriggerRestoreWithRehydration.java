
import com.azure.resourcemanager.dataprotection.models.AzureBackupRestoreWithRehydrationRequest;
import com.azure.resourcemanager.dataprotection.models.Datasource;
import com.azure.resourcemanager.dataprotection.models.DatasourceSet;
import com.azure.resourcemanager.dataprotection.models.RecoveryOption;
import com.azure.resourcemanager.dataprotection.models.RehydrationPriority;
import com.azure.resourcemanager.dataprotection.models.RestoreTargetInfo;
import com.azure.resourcemanager.dataprotection.models.SourceDataStoreType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for BackupInstances TriggerRestore.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01/BackupInstanceOperations/TriggerRestoreWithRehydration.json
     */
    /**
     * Sample code: Trigger Restore With Rehydration.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        triggerRestoreWithRehydration(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().triggerRestore("000pikumar", "PratikPrivatePreviewVault1", "testInstance1",
            new AzureBackupRestoreWithRehydrationRequest().withRestoreTargetInfo(new RestoreTargetInfo()
                .withRecoveryOption(RecoveryOption.FAIL_IF_EXISTS).withRestoreLocation("southeastasia")
                .withDatasourceInfo(new Datasource().withDatasourceType("OssDB").withObjectType("Datasource")
                    .withResourceId(
                        "/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb")
                    .withResourceLocation("").withResourceName("testdb")
                    .withResourceType("Microsoft.DBforPostgreSQL/servers/databases").withResourceUri(""))
                .withDatasourceSetInfo(new DatasourceSet().withDatasourceType("OssDB").withObjectType("DatasourceSet")
                    .withResourceId(
                        "/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest")
                    .withResourceLocation("").withResourceName("viveksipgtest")
                    .withResourceType("Microsoft.DBforPostgreSQL/servers").withResourceUri("")))
                .withSourceDataStoreType(SourceDataStoreType.VAULT_STORE)
                .withSourceResourceId(
                    "/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb")
                .withRecoveryPointId("hardcodedRP").withRehydrationPriority(RehydrationPriority.HIGH)
                .withRehydrationRetentionDuration("7D"),
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
