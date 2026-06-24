
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.compute.models.MigrateToVirtualMachineScaleSetInput;

/**
 * Samples for AvailabilitySets ValidateMigrationToVirtualMachineScaleSet.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/availabilitySetExamples/AvailabilitySet_ValidateMigrationToVirtualMachineScaleSet.json
     */
    /**
     * Sample code: AvailabilitySet_ValidateMigrationToVirtualMachineScaleSet.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void availabilitySetValidateMigrationToVirtualMachineScaleSet(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getAvailabilitySets().validateMigrationToVirtualMachineScaleSetWithResponse("rgcompute",
            "myAvailabilitySet",
            new MigrateToVirtualMachineScaleSetInput().withVirtualMachineScaleSetFlexible(new SubResource().withId(
                "/subscriptions/{subscription-id}/resourceGroups/rgcompute/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}")),
            com.azure.core.util.Context.NONE);
    }
}
