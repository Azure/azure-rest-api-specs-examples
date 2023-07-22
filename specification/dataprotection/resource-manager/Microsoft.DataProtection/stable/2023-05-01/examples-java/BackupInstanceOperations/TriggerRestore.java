import com.azure.resourcemanager.dataprotection.models.AzureBackupRecoveryPointBasedRestoreRequest;
import com.azure.resourcemanager.dataprotection.models.Datasource;
import com.azure.resourcemanager.dataprotection.models.DatasourceSet;
import com.azure.resourcemanager.dataprotection.models.RecoveryOption;
import com.azure.resourcemanager.dataprotection.models.RestoreTargetInfo;
import com.azure.resourcemanager.dataprotection.models.SecretStoreBasedAuthCredentials;
import com.azure.resourcemanager.dataprotection.models.SecretStoreResource;
import com.azure.resourcemanager.dataprotection.models.SecretStoreType;
import com.azure.resourcemanager.dataprotection.models.SourceDataStoreType;

/** Samples for BackupInstances TriggerRestore. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2023-05-01/examples/BackupInstanceOperations/TriggerRestore.json
     */
    /**
     * Sample code: Trigger Restore.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void triggerRestore(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager
            .backupInstances()
            .triggerRestore(
                "000pikumar",
                "PratikPrivatePreviewVault1",
                "testInstance1",
                new AzureBackupRecoveryPointBasedRestoreRequest()
                    .withRestoreTargetInfo(
                        new RestoreTargetInfo()
                            .withRecoveryOption(RecoveryOption.FAIL_IF_EXISTS)
                            .withRestoreLocation("southeastasia")
                            .withDatasourceInfo(
                                new Datasource()
                                    .withDatasourceType("Microsoft.DBforPostgreSQL/servers/databases")
                                    .withObjectType("Datasource")
                                    .withResourceId(
                                        "/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/targetdb")
                                    .withResourceLocation("")
                                    .withResourceName("targetdb")
                                    .withResourceType("Microsoft.DBforPostgreSQL/servers/databases")
                                    .withResourceUri(""))
                            .withDatasourceSetInfo(
                                new DatasourceSet()
                                    .withDatasourceType("Microsoft.DBforPostgreSQL/servers/databases")
                                    .withObjectType("DatasourceSet")
                                    .withResourceId(
                                        "/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest")
                                    .withResourceLocation("")
                                    .withResourceName("viveksipgtest")
                                    .withResourceType("Microsoft.DBforPostgreSQL/servers")
                                    .withResourceUri(""))
                            .withDatasourceAuthCredentials(
                                new SecretStoreBasedAuthCredentials()
                                    .withSecretStoreResource(
                                        new SecretStoreResource()
                                            .withUri("https://samplevault.vault.azure.net/secrets/credentials")
                                            .withSecretStoreType(SecretStoreType.AZURE_KEY_VAULT))))
                    .withSourceDataStoreType(SourceDataStoreType.VAULT_STORE)
                    .withSourceResourceId(
                        "/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb")
                    .withRecoveryPointId("hardcodedRP"),
                com.azure.core.util.Context.NONE);
    }
}
