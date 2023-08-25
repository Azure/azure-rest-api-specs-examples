/** Samples for VirtualMachines PerformMaintenance. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/virtualMachineExamples/VirtualMachine_PerformMaintenance_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachine_PerformMaintenance_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachinePerformMaintenanceMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .performMaintenance("rgcompute", "aaaaaaa", com.azure.core.util.Context.NONE);
    }
}
