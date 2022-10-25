import com.azure.core.util.Context;

/** Samples for ProtectionContainers Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/preview/2022-09-01-preview/examples/AzureWorkload/ProtectionContainers_Get.json
     */
    /**
     * Sample code: Get Protection Container Details.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getProtectionContainerDetails(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .protectionContainers()
            .getWithResponse("testVault", "testRg", "Azure", "VMAppContainer;Compute;testRG;testSQL", Context.NONE);
    }
}
