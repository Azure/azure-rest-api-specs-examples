
/**
 * Samples for VirtualMachineScaleSetVMs SimulateEviction.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_SimulateEviction.json
     */
    /**
     * Sample code: Simulate Eviction a virtual machine.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void simulateEvictionAVirtualMachine(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getVirtualMachineScaleSetVMs().simulateEvictionWithResponse("ResourceGroup",
            "VmScaleSetName", "InstanceId", com.azure.core.util.Context.NONE);
    }
}
