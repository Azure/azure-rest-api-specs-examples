
/**
 * Samples for DedicatedHostGroups GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-03-01/examples/
     * dedicatedHostExamples/DedicatedHostGroup_Get.json
     */
    /**
     * Sample code: Create a dedicated host group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createADedicatedHostGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDedicatedHostGroups().getByResourceGroupWithResponse(
            "myResourceGroup", "myDedicatedHostGroup", null, com.azure.core.util.Context.NONE);
    }
}
