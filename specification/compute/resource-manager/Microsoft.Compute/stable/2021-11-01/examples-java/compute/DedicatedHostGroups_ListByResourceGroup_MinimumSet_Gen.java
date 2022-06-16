import com.azure.core.util.Context;

/** Samples for DedicatedHostGroups ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/DedicatedHostGroups_ListByResourceGroup_MinimumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHostGroups_ListByResourceGroup_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dedicatedHostGroupsListByResourceGroupMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDedicatedHostGroups()
            .listByResourceGroup("rgcompute", Context.NONE);
    }
}
