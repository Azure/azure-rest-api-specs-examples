/** Samples for VirtualMachines InstanceView. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-03-01/examples/virtualMachineExamples/VirtualMachine_Get_InstanceView.json
     */
    /**
     * Sample code: Get Virtual Machine Instance View.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getVirtualMachineInstanceView(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getVirtualMachines()
            .instanceViewWithResponse("myResourceGroup", "myVM", com.azure.core.util.Context.NONE);
    }
}
