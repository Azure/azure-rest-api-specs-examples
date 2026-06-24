
/**
 * Samples for VirtualMachines InstanceView.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_Get_InstanceView_WithInterconnectBlock.json
     */
    /**
     * Sample code: Get instance view of a virtual machine associated with an Interconnect Block.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getInstanceViewOfAVirtualMachineAssociatedWithAnInterconnectBlock(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().instanceViewWithResponse("myResourceGroup", "myVM",
            com.azure.core.util.Context.NONE);
    }
}
