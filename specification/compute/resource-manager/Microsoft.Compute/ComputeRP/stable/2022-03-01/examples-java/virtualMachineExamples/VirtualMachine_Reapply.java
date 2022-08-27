import com.azure.core.util.Context;

/** Samples for VirtualMachines Reapply. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/virtualMachineExamples/VirtualMachine_Reapply.json
     */
    /**
     * Sample code: Reapply the state of a virtual machine.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void reapplyTheStateOfAVirtualMachine(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .reapply("ResourceGroup", "VMName", Context.NONE);
    }
}
