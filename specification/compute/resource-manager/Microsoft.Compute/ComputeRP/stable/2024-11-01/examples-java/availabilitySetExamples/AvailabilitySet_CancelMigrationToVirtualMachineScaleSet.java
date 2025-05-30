
/**
 * Samples for AvailabilitySets CancelMigrationToVirtualMachineScaleSet.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * availabilitySetExamples/AvailabilitySet_CancelMigrationToVirtualMachineScaleSet.json
     */
    /**
     * Sample code: AvailabilitySet_CancelMigrationToVirtualMachineScaleSet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        availabilitySetCancelMigrationToVirtualMachineScaleSet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getAvailabilitySets()
            .cancelMigrationToVirtualMachineScaleSetWithResponse("rgcompute", "myAvailabilitySet",
                com.azure.core.util.Context.NONE);
    }
}
