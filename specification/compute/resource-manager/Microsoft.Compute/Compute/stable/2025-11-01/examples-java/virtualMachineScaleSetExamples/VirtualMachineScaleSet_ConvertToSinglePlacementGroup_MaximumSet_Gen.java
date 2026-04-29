
import com.azure.resourcemanager.compute.models.VMScaleSetConvertToSinglePlacementGroupInput;

/**
 * Samples for VirtualMachineScaleSets ConvertToSinglePlacementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_ConvertToSinglePlacementGroup_MaximumSet_Gen.
     * json
     */
    /**
     * Sample code: VirtualMachineScaleSet_ConvertToSinglePlacementGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineScaleSetConvertToSinglePlacementGroupMaximumSetGen(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().convertToSinglePlacementGroupWithResponse("rgcompute",
            "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", new VMScaleSetConvertToSinglePlacementGroupInput()
                .withActivePlacementGroupId("aaaaaaaaaaaaaaaaaaaaaaaaaaa"),
            com.azure.core.util.Context.NONE);
    }
}
