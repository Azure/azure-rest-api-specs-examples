import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetVMs PerformMaintenance. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachineScaleSetVMs_PerformMaintenance_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSetVMs_PerformMaintenance_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetVMsPerformMaintenanceMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMs()
            .performMaintenance("rgcompute", "aaaaaaaaaaaaaa", "aaaaaaaaaaaa", Context.NONE);
    }
}
