import com.azure.core.util.Context;

/** Samples for DedicatedHostGroups GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/compute/GetADedicatedHostGroup.json
     */
    /**
     * Sample code: Create a dedicated host group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createADedicatedHostGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDedicatedHostGroups()
            .getByResourceGroupWithResponse("myResourceGroup", "myDedicatedHostGroup", null, Context.NONE);
    }
}
