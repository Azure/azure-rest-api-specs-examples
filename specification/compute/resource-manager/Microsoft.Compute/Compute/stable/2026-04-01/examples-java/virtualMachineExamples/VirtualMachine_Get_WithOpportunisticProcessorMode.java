
/**
 * Samples for VirtualMachines GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/virtualMachineExamples/VirtualMachine_Get_WithOpportunisticProcessorMode.json
     */
    /**
     * Sample code: Get a virtual machine with Opportunistic Processor Mode.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getAVirtualMachineWithOpportunisticProcessorMode(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().getByResourceGroupWithResponse("myResourceGroup", "myVM", null,
            com.azure.core.util.Context.NONE);
    }
}
