
import com.azure.resourcemanager.compute.models.ExpandTypesForGetVMScaleSets;

/**
 * Samples for VirtualMachineScaleSets GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-03-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Get_WithDiskControllerType.json
     */
    /**
     * Sample code: Get VM scale set VM with Disk Controller Type.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getVMScaleSetVMWithDiskControllerType(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().getByResourceGroupWithResponse("myResourceGroup",
            "myVirtualMachineScaleSet", ExpandTypesForGetVMScaleSets.USER_DATA, com.azure.core.util.Context.NONE);
    }
}
