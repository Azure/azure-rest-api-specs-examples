import com.azure.core.util.Context;

/** Samples for VirtualMachines Generalize. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/GeneralizeVirtualMachine.json
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
            .generalizeWithResponse("myResourceGroup", "myVMName", Context.NONE);
    }
}
