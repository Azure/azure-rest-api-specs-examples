
import com.azure.resourcemanager.dataprotection.models.AdHocBackupRuleOptions;
import com.azure.resourcemanager.dataprotection.models.AdhocBackupTriggerOption;
import com.azure.resourcemanager.dataprotection.models.TriggerBackupRequest;

/**
 * Samples for BackupInstances AdhocBackup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/
     * BackupInstanceOperations/TriggerBackup.json
     */
    /**
     * Sample code: Trigger Adhoc Backup.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void triggerAdhocBackup(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.backupInstances().adhocBackup("000pikumar", "PratikPrivatePreviewVault1", "testInstance1",
            new TriggerBackupRequest().withBackupRuleOptions(new AdHocBackupRuleOptions().withRuleName("BackupWeekly")
                .withTriggerOption(new AdhocBackupTriggerOption().withRetentionTagOverride("yearly"))),
            com.azure.core.util.Context.NONE);
    }
}
