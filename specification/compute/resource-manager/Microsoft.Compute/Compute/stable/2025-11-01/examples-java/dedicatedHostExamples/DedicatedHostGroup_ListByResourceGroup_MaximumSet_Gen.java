
/**
 * Samples for DedicatedHostGroups ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/dedicatedHostExamples/DedicatedHostGroup_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHostGroup_ListByResourceGroup_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        dedicatedHostGroupListByResourceGroupMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDedicatedHostGroups().listByResourceGroup("rgcompute",
            com.azure.core.util.Context.NONE);
    }
}
