
import com.azure.resourcemanager.dataprotection.models.SyncBackupInstanceRequest;
import com.azure.resourcemanager.dataprotection.models.SyncType;

/**
 * Samples for BackupInstances SyncBackupInstance.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/
     * BackupInstanceOperations/SyncBackupInstance.json
     */
    /**
     * Sample code: Sync BackupInstance.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void syncBackupInstance(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().syncBackupInstance("testrg", "testvault", "testbi",
            new SyncBackupInstanceRequest().withSyncType(SyncType.DEFAULT), com.azure.core.util.Context.NONE);
    }
}
