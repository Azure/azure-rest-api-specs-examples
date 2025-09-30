
import com.azure.resourcemanager.compute.models.VMScaleSetConvertToSinglePlacementGroupInput;

/**
 * Samples for VirtualMachineScaleSets ConvertToSinglePlacementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2025-04-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSet_ConvertToSinglePlacementGroup_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSet_ConvertToSinglePlacementGroup_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetConvertToSinglePlacementGroupMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSets()
            .convertToSinglePlacementGroupWithResponse("rgcompute", "aaaaaaaaaaaaa",
                new VMScaleSetConvertToSinglePlacementGroupInput(), com.azure.core.util.Context.NONE);
    }
}
