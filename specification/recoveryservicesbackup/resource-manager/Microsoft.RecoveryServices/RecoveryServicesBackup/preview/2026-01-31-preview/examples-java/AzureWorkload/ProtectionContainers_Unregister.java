
/**
 * Samples for ProtectionContainers Unregister.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/AzureWorkload/ProtectionContainers_Unregister.json
     */
    /**
     * Sample code: Unregister Protection Container.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void unregisterProtectionContainer(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.protectionContainers().unregisterWithResponse("testVault", "testRg", "Azure",
            "storagecontainer;Storage;test-rg;teststorage", com.azure.core.util.Context.NONE);
    }
}
