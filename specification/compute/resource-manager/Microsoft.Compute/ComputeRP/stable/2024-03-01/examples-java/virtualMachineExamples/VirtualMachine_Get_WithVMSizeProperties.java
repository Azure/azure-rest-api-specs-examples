
import com.azure.resourcemanager.compute.models.InstanceViewTypes;

/**
 * Samples for VirtualMachines GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * virtualMachineExamples/VirtualMachine_Get_WithVMSizeProperties.json
     */
    /**
     * Sample code: Get a virtual machine with VM Size Properties.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAVirtualMachineWithVMSizeProperties(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachines()
            .getByResourceGroupWithResponse("myResourceGroup", "myVM", null, com.azure.core.util.Context.NONE);
    }

    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * virtualMachineExamples/VirtualMachine_Get_WithDiskControllerType.json
     */
    /**
     * Sample code: Get a virtual machine with Disk Controller Type Properties.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getAVirtualMachineWithDiskControllerTypeProperties(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachines().getByResourceGroupWithResponse(
            "myResourceGroup", "myVM", InstanceViewTypes.USER_DATA, com.azure.core.util.Context.NONE);
    }

    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * virtualMachineExamples/VirtualMachine_Get.json
     */
    /**
     * Sample code: Get a Virtual Machine.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAVirtualMachine(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachines().getByResourceGroupWithResponse(
            "myResourceGroup", "myVM", InstanceViewTypes.USER_DATA, com.azure.core.util.Context.NONE);
    }
}
