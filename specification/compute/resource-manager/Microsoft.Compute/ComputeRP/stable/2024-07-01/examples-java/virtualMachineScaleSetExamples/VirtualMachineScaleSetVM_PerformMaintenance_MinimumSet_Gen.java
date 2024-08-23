
/**
 * Samples for VirtualMachineScaleSetVMs PerformMaintenance.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_PerformMaintenance_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVM_PerformMaintenance_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        virtualMachineScaleSetVMPerformMaintenanceMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetVMs().performMaintenance("rgcompute",
            "aaaaaaaaaa", "aaaa", com.azure.core.util.Context.NONE);
    }
}
