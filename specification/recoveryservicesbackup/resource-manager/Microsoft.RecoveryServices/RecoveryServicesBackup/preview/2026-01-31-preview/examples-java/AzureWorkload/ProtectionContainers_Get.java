
/**
 * Samples for ProtectionContainers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-31-preview/AzureWorkload/ProtectionContainers_Get.json
     */
    /**
     * Sample code: Get Protection Container Details.
     * 
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getProtectionContainerDetails(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager.protectionContainers().getWithResponse("testVault", "testRg", "Azure",
            "VMAppContainer;Compute;testRG;testSQL", com.azure.core.util.Context.NONE);
    }
}
