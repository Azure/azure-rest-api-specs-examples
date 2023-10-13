/** Samples for ProtectionIntent Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2023-04-01/examples/AzureWorkload/BackupProtectionIntent_Get.json
     */
    /**
     * Sample code: Get ProtectionIntent for an item.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getProtectionIntentForAnItem(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .protectionIntents()
            .getWithResponse(
                "myVault", "myRG", "Azure", "249D9B07-D2EF-4202-AA64-65F35418564E", com.azure.core.util.Context.NONE);
    }
}
