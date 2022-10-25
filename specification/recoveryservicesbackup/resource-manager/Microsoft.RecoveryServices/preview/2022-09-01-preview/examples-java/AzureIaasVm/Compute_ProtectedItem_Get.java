import com.azure.core.util.Context;

/** Samples for ProtectedItems Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesbackup/resource-manager/Microsoft.RecoveryServices/preview/2022-09-01-preview/examples/AzureIaasVm/Compute_ProtectedItem_Get.json
     */
    /**
     * Sample code: Get Protected Virtual Machine Details.
     *
     * @param manager Entry point to RecoveryServicesBackupManager.
     */
    public static void getProtectedVirtualMachineDetails(
        com.azure.resourcemanager.recoveryservicesbackup.RecoveryServicesBackupManager manager) {
        manager
            .protectedItems()
            .getWithResponse(
                "PySDKBackupTestRsVault",
                "PythonSDKBackupTestRg",
                "Azure",
                "iaasvmcontainer;iaasvmcontainerv2;iaasvm-rg;iaasvm-1",
                "vm;iaasvmcontainerv2;iaasvm-rg;iaasvm-1",
                null,
                Context.NONE);
    }
}
