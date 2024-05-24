
/**
 * Samples for ProtectionIntent Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2024-04-01/examples/
     * AzureWorkload/BackupProtectionIntent_Delete.json
     */
    /**
     * Sample code: Delete Protection intent from item.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void deleteProtectionIntentFromItem(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.protectionIntents().deleteWithResponse("myVault", "myRG", "Azure",
            "249D9B07-D2EF-4202-AA64-65F35418564E", com.azure.core.util.Context.NONE);
    }
}
