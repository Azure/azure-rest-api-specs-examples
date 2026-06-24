
/**
 * Samples for VirtualMachines SimulateEviction.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/virtualMachineExamples/VirtualMachine_SimulateEviction.json
     */
    /**
     * Sample code: Simulate Eviction a virtual machine.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void simulateEvictionAVirtualMachine(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachines().simulateEvictionWithResponse("ResourceGroup", "VMName",
            com.azure.core.util.Context.NONE);
    }
}
