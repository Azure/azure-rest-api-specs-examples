
/**
 * Samples for ResourceGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01/ForceDeleteVMsInResourceGroup.json
     */
    /**
     * Sample code: Force delete all the Virtual Machines in a resource group.
     * 
     * @param manager Entry point to ResourceManager.
     */
    public static void
        forceDeleteAllTheVirtualMachinesInAResourceGroup(com.azure.resourcemanager.resources.ResourceManager manager) {
        manager.serviceClient().getResourceGroups().delete("my-resource-group", "Microsoft.Compute/virtualMachines",
            com.azure.core.util.Context.NONE);
    }
}
