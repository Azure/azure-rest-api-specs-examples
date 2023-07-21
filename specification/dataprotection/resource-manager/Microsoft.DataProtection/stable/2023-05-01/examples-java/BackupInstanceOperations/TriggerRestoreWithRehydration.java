import com.azure.resourcemanager.dataprotection.models.AzureBackupRestoreWithRehydrationRequest;
import com.azure.resourcemanager.dataprotection.models.Datasource;
import com.azure.resourcemanager.dataprotection.models.DatasourceSet;
import com.azure.resourcemanager.dataprotection.models.RecoveryOption;
import com.azure.resourcemanager.dataprotection.models.RehydrationPriority;
import com.azure.resourcemanager.dataprotection.models.RestoreTargetInfo;
import com.azure.resourcemanager.dataprotection.models.SourceDataStoreType;

/** Samples for BackupInstances TriggerRestore. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-05-01/examples/BackupInstanceOperations/TriggerRestoreWithRehydration.json
     */
    /**
     * Sample code: Trigger Restore With Rehydration.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void triggerRestoreWithRehydration(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager
            .backupInstances()
            .triggerRestore(
                "000pikumar",
                "PratikPrivatePreviewVault1",
                "testInstance1",
                new AzureBackupRestoreWithRehydrationRequest()
                    .withRestoreTargetInfo(
                        new RestoreTargetInfo()
                            .withRecoveryOption(RecoveryOption.FAIL_IF_EXISTS)
                            .withRestoreLocation("southeastasia")
                            .withDatasourceInfo(
                                new Datasource()
                                    .withDatasourceType("OssDB")
                                    .withObjectType("Datasource")
                                    .withResourceId(
                                        "/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb")
                                    .withResourceLocation("")
                                    .withResourceName("testdb")
                                    .withResourceType("Microsoft.DBforPostgreSQL/servers/databases")
                                    .withResourceUri(""))
                            .withDatasourceSetInfo(
                                new DatasourceSet()
                                    .withDatasourceType("OssDB")
                                    .withObjectType("DatasourceSet")
                                    .withResourceId(
                                        "/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest")
                                    .withResourceLocation("")
                                    .withResourceName("viveksipgtest")
                                    .withResourceType("Microsoft.DBforPostgreSQL/servers")
                                    .withResourceUri("")))
                    .withSourceDataStoreType(SourceDataStoreType.VAULT_STORE)
                    .withSourceResourceId(
                        "/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb")
                    .withRecoveryPointId("hardcodedRP")
                    .withRehydrationPriority(RehydrationPriority.HIGH)
                    .withRehydrationRetentionDuration("7D"),
                com.azure.core.util.Context.NONE);
    }
}
