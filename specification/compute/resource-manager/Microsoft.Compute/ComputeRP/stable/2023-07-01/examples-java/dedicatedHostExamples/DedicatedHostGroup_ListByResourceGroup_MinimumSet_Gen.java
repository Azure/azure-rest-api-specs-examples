/** Samples for DedicatedHostGroups ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/dedicatedHostExamples/DedicatedHostGroup_ListByResourceGroup_MinimumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHostGroup_ListByResourceGroup_MinimumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dedicatedHostGroupListByResourceGroupMinimumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDedicatedHostGroups()
            .listByResourceGroup("rgcompute", com.azure.core.util.Context.NONE);
    }
}
