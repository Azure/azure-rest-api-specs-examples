import com.azure.core.util.Context;

/** Samples for VirtualMachines PerformMaintenance. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/VirtualMachines_PerformMaintenance_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachines_PerformMaintenance_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachinesPerformMaintenanceMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .performMaintenance("rgcompute", "aaaaaaaaaa", Context.NONE);
    }
}
