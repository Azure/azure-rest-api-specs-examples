
import com.azure.resourcemanager.compute.models.VMScaleSetConvertToSinglePlacementGroupInput;

/**
 * Samples for VirtualMachineScaleSets ConvertToSinglePlacementGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_ConvertToSinglePlacementGroup_MinimumSet_Gen.
     * json
     */
    /**
     * Sample code: VirtualMachineScaleSet_ConvertToSinglePlacementGroup_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void virtualMachineScaleSetConvertToSinglePlacementGroupMinimumSetGen(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().convertToSinglePlacementGroupWithResponse("rgcompute",
            "aaaaaaaaaaaaa", new VMScaleSetConvertToSinglePlacementGroupInput(), com.azure.core.util.Context.NONE);
    }
}
