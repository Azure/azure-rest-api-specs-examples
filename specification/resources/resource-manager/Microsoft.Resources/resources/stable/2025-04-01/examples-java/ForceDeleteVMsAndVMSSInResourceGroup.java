
/**
 * Samples for ResourceGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01/ForceDeleteVMsAndVMSSInResourceGroup.json
     */
    /**
     * Sample code: Force delete all the Virtual Machines and Virtual Machine Scale Sets in a resource group.
     * 
     * @param manager Entry point to ResourceManager.
     */
    public static void forceDeleteAllTheVirtualMachinesAndVirtualMachineScaleSetsInAResourceGroup(
        com.azure.resourcemanager.resources.ResourceManager manager) {
        manager.serviceClient().getResourceGroups().delete("my-resource-group",
            "Microsoft.Compute/virtualMachines,Microsoft.Compute/virtualMachineScaleSets",
            com.azure.core.util.Context.NONE);
    }
}
