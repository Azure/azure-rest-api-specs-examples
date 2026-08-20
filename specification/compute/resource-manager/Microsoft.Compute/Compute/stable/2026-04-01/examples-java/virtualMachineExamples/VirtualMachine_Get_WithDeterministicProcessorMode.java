
/**
 * Samples for VirtualMachines GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/virtualMachineExamples/VirtualMachine_Get_WithDeterministicProcessorMode.json
     */
    /**
     * Sample code: Get a virtual machine with Deterministic Processor Mode.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getAVirtualMachineWithDeterministicProcessorMode(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().getByResourceGroupWithResponse("myResourceGroup", "myVM", null,
            com.azure.core.util.Context.NONE);
    }
}
