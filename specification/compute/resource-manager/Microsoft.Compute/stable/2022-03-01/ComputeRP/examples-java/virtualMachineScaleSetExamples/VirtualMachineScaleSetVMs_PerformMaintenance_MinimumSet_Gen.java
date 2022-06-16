import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetVMs PerformMaintenance. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSetVMs_PerformMaintenance_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVMs_PerformMaintenance_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetVMsPerformMaintenanceMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMs()
            .performMaintenance("rgcompute", "aaaaaaaaaa", "aaaa", Context.NONE);
    }
}
