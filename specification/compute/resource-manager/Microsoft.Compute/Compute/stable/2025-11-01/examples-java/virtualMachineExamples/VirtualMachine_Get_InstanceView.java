
/**
 * Samples for VirtualMachines InstanceView.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachine_Get_InstanceView.json
     */
    /**
     * Sample code: Get Virtual Machine Instance View.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getVirtualMachineInstanceView(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().instanceViewWithResponse("myResourceGroup", "myVM",
            com.azure.core.util.Context.NONE);
    }
}
