
import com.azure.resourcemanager.compute.models.InstanceViewTypes;

/**
 * Samples for VirtualMachines GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachine_Get_WithDiskControllerType.json
     */
    /**
     * Sample code: Get a virtual machine with Disk Controller Type Properties.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getAVirtualMachineWithDiskControllerTypeProperties(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().getByResourceGroupWithResponse("myResourceGroup", "myVM",
            InstanceViewTypes.USER_DATA, com.azure.core.util.Context.NONE);
    }
}
