/** Samples for VirtualMachines Generalize. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/virtualMachineExamples/VirtualMachine_Generalize.json
     */
    /**
     * Sample code: Generalize a Virtual Machine.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void generalizeAVirtualMachine(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .generalizeWithResponse("myResourceGroup", "myVMName", com.azure.core.util.Context.NONE);
    }
}
