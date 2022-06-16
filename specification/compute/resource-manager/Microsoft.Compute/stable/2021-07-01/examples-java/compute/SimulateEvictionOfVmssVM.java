import com.azure.core.util.Context;

/** Samples for VirtualMachineScaleSetVMs SimulateEviction. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/compute/SimulateEvictionOfVmssVM.json
     */
    /**
     * Sample code: Simulate Eviction a virtual machine.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void simulateEvictionAVirtualMachine(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachineScaleSetVMs()
            .simulateEvictionWithResponse("ResourceGroup", "VmScaleSetName", "InstanceId", Context.NONE);
    }
}
