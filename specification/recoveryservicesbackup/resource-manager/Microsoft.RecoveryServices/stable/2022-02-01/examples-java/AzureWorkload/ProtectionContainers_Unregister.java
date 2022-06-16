import com.azure.core.util.Context;

/** Samples for ProtectionContainers Unregister. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/stable/2022-02-01/examples/AzureWorkload/ProtectionContainers_Unregister.json
     */
    /**
     * Sample code: Unregister Protection Container.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void unregisterProtectionContainer(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .protectionContainers()
            .unregisterWithResponse(
                "testVault", "testRg", "Azure", "storagecontainer;Storage;test-rg;teststorage", Context.NONE);
    }
}
