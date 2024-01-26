
/**
 * Samples for VirtualMachineScaleSetVMs SimulateEviction.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/
     * virtualMachineScaleSetExamples/VirtualMachineScaleSetVM_SimulateEviction.json
     */
    /**
     * Sample code: Simulate Eviction a virtual machine.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void simulateEvictionAVirtualMachine(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getVirtualMachineScaleSetVMs().simulateEvictionWithResponse(
            "ResourceGroup", "VmScaleSetName", "InstanceId", com.azure.core.util.Context.NONE);
    }
}
