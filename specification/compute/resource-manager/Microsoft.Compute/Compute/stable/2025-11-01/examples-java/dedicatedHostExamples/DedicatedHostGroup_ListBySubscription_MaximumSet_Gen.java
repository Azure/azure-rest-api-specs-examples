
/**
 * Samples for DedicatedHostGroups List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/dedicatedHostExamples/DedicatedHostGroup_ListBySubscription_MaximumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHostGroup_ListBySubscription_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        dedicatedHostGroupListBySubscriptionMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDedicatedHostGroups().list(com.azure.core.util.Context.NONE);
    }
}
