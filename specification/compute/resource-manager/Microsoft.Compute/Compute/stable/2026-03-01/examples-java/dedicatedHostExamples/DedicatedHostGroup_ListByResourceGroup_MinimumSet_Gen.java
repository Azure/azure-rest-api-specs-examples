
/**
 * Samples for DedicatedHostGroups ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/dedicatedHostExamples/DedicatedHostGroup_ListByResourceGroup_MinimumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHostGroup_ListByResourceGroup_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        dedicatedHostGroupListByResourceGroupMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDedicatedHostGroups().listByResourceGroup("rgcompute",
            com.azure.core.util.Context.NONE);
    }
}
