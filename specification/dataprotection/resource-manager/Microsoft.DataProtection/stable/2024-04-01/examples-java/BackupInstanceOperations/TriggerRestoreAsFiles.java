
import com.azure.resourcemanager.dataprotection.models.AzureBackupRecoveryPointBasedRestoreRequest;
import com.azure.resourcemanager.dataprotection.models.RecoveryOption;
import com.azure.resourcemanager.dataprotection.models.RestoreFilesTargetInfo;
import com.azure.resourcemanager.dataprotection.models.RestoreTargetLocationType;
import com.azure.resourcemanager.dataprotection.models.SourceDataStoreType;
import com.azure.resourcemanager.dataprotection.models.TargetDetails;

/**
 * Samples for BackupInstances TriggerRestore.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/
     * BackupInstanceOperations/TriggerRestoreAsFiles.json
     */
    /**
     * Sample code: Trigger Restore As Files.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void triggerRestoreAsFiles(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().triggerRestore("000pikumar", "PrivatePreviewVault1", "testInstance1",
            new AzureBackupRecoveryPointBasedRestoreRequest()
                .withRestoreTargetInfo(new RestoreFilesTargetInfo().withRecoveryOption(RecoveryOption.FAIL_IF_EXISTS)
                    .withRestoreLocation("southeastasia")
                    .withTargetDetails(new TargetDetails().withFilePrefix("restoredblob")
                        .withRestoreTargetLocationType(RestoreTargetLocationType.AZURE_BLOBS)
                        .withUrl("https://teststorage.blob.core.windows.net/restoretest")))
                .withSourceDataStoreType(SourceDataStoreType.VAULT_STORE)
                .withSourceResourceId(
                    "/subscriptions/f75d8d8b-6735-4697-82e1-1a7a3ff0d5d4/resourceGroups/viveksipgtest/providers/Microsoft.DBforPostgreSQL/servers/viveksipgtest/databases/testdb")
                .withRecoveryPointId("hardcodedRP"),
            com.azure.core.util.Context.NONE);
    }
}
