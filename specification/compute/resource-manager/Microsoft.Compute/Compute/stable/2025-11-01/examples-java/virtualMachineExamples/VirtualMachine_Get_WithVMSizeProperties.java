
/**
 * Samples for VirtualMachines GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachine_Get_WithVMSizeProperties.json
     */
    /**
     * Sample code: Get a virtual machine with VM Size Properties.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getAVirtualMachineWithVMSizeProperties(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().getByResourceGroupWithResponse("myResourceGroup", "myVM", null,
            com.azure.core.util.Context.NONE);
    }
}
