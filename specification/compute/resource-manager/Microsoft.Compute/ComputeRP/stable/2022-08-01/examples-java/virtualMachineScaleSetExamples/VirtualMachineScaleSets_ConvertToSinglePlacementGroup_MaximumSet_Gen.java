import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.VMScaleSetConvertToSinglePlacementGroupInput;

/** Samples for VirtualMachineScaleSets ConvertToSinglePlacementGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSets_ConvertToSinglePlacementGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: VirtualMachineScaleSets_ConvertToSinglePlacementGroup_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void virtualMachineScaleSetsConvertToSinglePlacementGroupMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .convertToSinglePlacementGroupWithResponse(
                "rgcompute",
                "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
                new VMScaleSetConvertToSinglePlacementGroupInput()
                    .withActivePlacementGroupId("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
                Context.NONE);
    }
}
