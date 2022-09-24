import com.azure.core.util.Context;

/** Samples for VirtualMachines PerformMaintenance. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineExamples/VirtualMachines_PerformMaintenance_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachines_PerformMaintenance_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachinesPerformMaintenanceMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .performMaintenance("rgcompute", "aaaaaaa", Context.NONE);
    }
}
