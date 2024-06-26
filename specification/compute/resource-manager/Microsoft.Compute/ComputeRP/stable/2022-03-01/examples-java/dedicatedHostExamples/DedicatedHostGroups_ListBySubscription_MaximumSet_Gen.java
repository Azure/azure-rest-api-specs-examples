import com.azure.core.util.Context;

/** Samples for DedicatedHostGroups List. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-03-01/examples/dedicatedHostExamples/DedicatedHostGroups_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHostGroups_ListBySubscription_MaximumSet_Gen.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void dedicatedHostGroupsListBySubscriptionMaximumSetGen(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDedicatedHostGroups().list(Context.NONE);
    }
}
