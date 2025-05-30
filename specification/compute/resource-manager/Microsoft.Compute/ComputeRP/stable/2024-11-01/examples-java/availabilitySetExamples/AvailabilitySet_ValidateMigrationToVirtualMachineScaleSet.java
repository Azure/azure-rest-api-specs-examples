
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.compute.models.MigrateToVirtualMachineScaleSetInput;

/**
 * Samples for AvailabilitySets ValidateMigrationToVirtualMachineScaleSet.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * availabilitySetExamples/AvailabilitySet_ValidateMigrationToVirtualMachineScaleSet.json
     */
    /**
     * Sample code: AvailabilitySet_ValidateMigrationToVirtualMachineScaleSet.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        availabilitySetValidateMigrationToVirtualMachineScaleSet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getAvailabilitySets()
            .validateMigrationToVirtualMachineScaleSetWithResponse("rgcompute", "myAvailabilitySet",
                new MigrateToVirtualMachineScaleSetInput().withVirtualMachineScaleSetFlexible(new SubResource().withId(
                    "/subscriptions/{subscription-id}/resourceGroups/rgcompute/providers/Microsoft.Compute/virtualMachineScaleSets/{vmss-name}")),
                com.azure.core.util.Context.NONE);
    }
}
