
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.compute.models.MigrateToVirtualMachineScaleSetInput;

/**
 * Samples for AvailabilitySets StartMigrationToVirtualMachineScaleSet.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/availabilitySetExamples/AvailabilitySet_StartMigrationToVirtualMachineScaleSet.json
     */
    /**
     * Sample code: AvailabilitySet_StartMigrationToVirtualMachineScaleSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void availabilitySetStartMigrationToVirtualMachineScaleSetGen(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getAvailabilitySets().startMigrationToVirtualMachineScaleSetWithResponse("rgcompute",
            "myAvailabilitySet",
            new MigrateToVirtualMachineScaleSetInput().withVirtualMachineScaleSetFlexible(new SubResource().withId(
                "/subscriptions/{subscription-id}/resourceGroups/rgcompute/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}")),
            com.azure.core.util.Context.NONE);
    }
}
