
/**
 * Samples for ProtectionContainers Inquire.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/AzureStorage/ProtectionContainers_Inquire.json
     */
    /**
     * Sample code: Inquire Azure Storage Protection Containers.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void inquireAzureStorageProtectionContainers(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.protectionContainers().inquireWithResponse("testvault", "test-rg", "Azure",
            "storagecontainer;Storage;test-rg;teststorage", null, com.azure.core.util.Context.NONE);
    }
}
