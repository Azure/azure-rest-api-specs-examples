
import com.azure.resourcemanager.dataprotection.models.AzureBackupRehydrationRequest;
import com.azure.resourcemanager.dataprotection.models.RehydrationPriority;

/**
 * Samples for BackupInstances TriggerRehydrate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/BackupInstanceOperations/TriggerRehydrate.json
     */
    /**
     * Sample code: Trigger Rehydrate.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void triggerRehydrate(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().triggerRehydrate("000pikumar", "PratikPrivatePreviewVault1", "testInstance1",
            new AzureBackupRehydrationRequest().withRecoveryPointId("hardcodedRP")
                .withRehydrationPriority(RehydrationPriority.HIGH).withRehydrationRetentionDuration("7D"),
            com.azure.core.util.Context.NONE);
    }
}
