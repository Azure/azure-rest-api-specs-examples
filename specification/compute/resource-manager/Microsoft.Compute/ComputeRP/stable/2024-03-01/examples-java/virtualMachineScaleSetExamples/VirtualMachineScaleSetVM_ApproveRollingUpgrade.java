
/**
 * Samples for VirtualMachineScaleSetVMs ApproveRollingUpgrade.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_ApproveRollingUpgrade.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_ApproveRollingUpgrade.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineScaleSetVMApproveRollingUpgrade(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetVMs().approveRollingUpgrade(
            "rgcompute", "vmssToApproveRollingUpgradeOn", "0123", com.azure.core.util.Context.NONE);
    }
}
