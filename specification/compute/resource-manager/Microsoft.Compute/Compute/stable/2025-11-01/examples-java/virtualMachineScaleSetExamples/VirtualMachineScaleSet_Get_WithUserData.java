
import com.azure.resourcemanager.compute.models.ExpandTypesForGetVMScaleSets;

/**
 * Samples for VirtualMachineScaleSets GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Get_WithUserData.json
     */
    /**
     * Sample code: Get a virtual machine scale set with UserData.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getAVirtualMachineScaleSetWithUserData(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSets().getByResourceGroupWithResponse("myResourceGroup",
            "myVirtualMachineScaleSet", ExpandTypesForGetVMScaleSets.USER_DATA, com.azure.core.util.Context.NONE);
    }
}
