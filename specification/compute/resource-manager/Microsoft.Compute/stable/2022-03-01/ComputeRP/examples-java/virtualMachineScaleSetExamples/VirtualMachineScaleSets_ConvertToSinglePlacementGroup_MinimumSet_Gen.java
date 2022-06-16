import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.VMScaleSetConvertToSinglePlacementGroupInput;

/** Samples for VirtualMachineScaleSets ConvertToSinglePlacementGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSets_ConvertToSinglePlacementGroup_MinimumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSets_ConvertToSinglePlacementGroup_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetsConvertToSinglePlacementGroupMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .convertToSinglePlacementGroupWithResponse(
                "rgcompute", "aaaaaaaaaaaaa", new VMScaleSetConvertToSinglePlacementGroupInput(), Context.NONE);
    }
}
