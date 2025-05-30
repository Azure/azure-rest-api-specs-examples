
/**
 * Samples for DedicatedHostGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/
     * dedicatedHostExamples/DedicatedHostGroup_ListBySubscription_MinimumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHostGroup_ListBySubscription_MinimumSet_Gen.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        dedicatedHostGroupListBySubscriptionMinimumSetGen(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDedicatedHostGroups()
            .list(com.azure.core.util.Context.NONE);
    }
}
