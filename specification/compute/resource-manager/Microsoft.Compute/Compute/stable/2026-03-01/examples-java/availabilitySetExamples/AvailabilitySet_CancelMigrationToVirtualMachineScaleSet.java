
/**
 * Samples for AvailabilitySets CancelMigrationToVirtualMachineScaleSet.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/availabilitySetExamples/AvailabilitySet_CancelMigrationToVirtualMachineScaleSet.json
     */
    /**
     * Sample code: AvailabilitySet_CancelMigrationToVirtualMachineScaleSet.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void availabilitySetCancelMigrationToVirtualMachineScaleSet(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getAvailabilitySets().cancelMigrationToVirtualMachineScaleSetWithResponse("rgcompute",
            "myAvailabilitySet", com.azure.core.util.Context.NONE);
    }
}
