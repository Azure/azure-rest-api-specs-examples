
import com.azure.resourcemanager.compute.models.InstanceViewTypes;

/**
 * Samples for VirtualMachines GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
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
}
