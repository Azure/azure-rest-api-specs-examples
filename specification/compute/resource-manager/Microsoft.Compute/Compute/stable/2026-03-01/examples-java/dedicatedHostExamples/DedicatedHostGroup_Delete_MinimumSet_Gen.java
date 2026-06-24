
/**
 * Samples for DedicatedHostGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/dedicatedHostExamples/DedicatedHostGroup_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHostGroup_Delete_MinimumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void dedicatedHostGroupDeleteMinimumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDedicatedHostGroups().deleteWithResponse("rgcompute", "aaaa",
            com.azure.core.util.Context.NONE);
    }
}
