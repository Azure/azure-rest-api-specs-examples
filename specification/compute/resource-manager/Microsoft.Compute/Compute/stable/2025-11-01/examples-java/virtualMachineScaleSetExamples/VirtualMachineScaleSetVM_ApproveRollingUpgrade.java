
/**
 * Samples for VirtualMachineScaleSetVMs ApproveRollingUpgrade.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_ApproveRollingUpgrade.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_ApproveRollingUpgrade.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        virtualMachineScaleSetVMApproveRollingUpgrade(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().approveRollingUpgrade("rgcompute",
            "vmssToApproveRollingUpgradeOn", "0123", com.azure.core.util.Context.NONE);
    }
}
