import com.azure.resourcemanager.compute.models.ExpandTypesForGetVMScaleSets;

/** Samples for VirtualMachineScaleSets GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Get_WithDiskControllerType.json
     */
    /**
     * Sample code: Get VM scale set VM with Disk Controller Type.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVMScaleSetVMWithDiskControllerType(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSets()
            .getByResourceGroupWithResponse(
                "myResourceGroup",
                "myVirtualMachineScaleSet",
                ExpandTypesForGetVMScaleSets.USER_DATA,
                com.azure.core.util.Context.NONE);
    }
}
