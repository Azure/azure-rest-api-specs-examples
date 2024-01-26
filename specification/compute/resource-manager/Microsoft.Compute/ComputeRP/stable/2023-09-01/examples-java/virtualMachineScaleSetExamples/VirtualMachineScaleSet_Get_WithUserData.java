
import com.azure.resourcemanager.compute.models.ExpandTypesForGetVMScaleSets;

/**
 * Samples for VirtualMachineScaleSets GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSet_Get_WithUserData.json
     */
    /**
     * Sample code: Get a virtual machine scale set with UserData.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAVirtualMachineScaleSetWithUserData(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSets().getByResourceGroupWithResponse(
            "myResourceGroup", "myVirtualMachineScaleSet", ExpandTypesForGetVMScaleSets.USER_DATA,
            com.azure.core.util.Context.NONE);
    }
}
