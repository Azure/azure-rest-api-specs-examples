
/**
 * Samples for ResourceGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2024-03-01/examples/
     * ForceDeleteVMsAndVMSSInResourceGroup.json
     */
    /**
     * Sample code: Force delete all the Virtual Machines and Virtual Machine Scale Sets in a resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void forceDeleteAllTheVirtualMachinesAndVirtualMachineScaleSetsInAResourceGroup(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().serviceClient().getResourceGroups().delete("my-resource-group",
            "Microsoft.Compute/virtualMachines,Microsoft.Compute/virtualMachineScaleSets",
            com.azure.core.util.Context.NONE);
    }
}
