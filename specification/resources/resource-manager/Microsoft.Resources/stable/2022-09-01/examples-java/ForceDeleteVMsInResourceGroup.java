
/** Samples for ResourceGroups Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2022-09-01/examples/
     * ForceDeleteVMsInResourceGroup.json
     */
    /**
     * Sample code: Force delete all the Virtual Machines in a resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        forceDeleteAllTheVirtualMachinesInAResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().serviceClient().getResourceGroups().delete("my-resource-group",
            "Microsoft.Compute/virtualMachines", com.azure.core.util.Context.NONE);
    }
}
