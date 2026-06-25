
/**
 * Samples for DedicatedHostGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/dedicatedHostExamples/DedicatedHostGroup_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: DedicatedHostGroup_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void dedicatedHostGroupDeleteMaximumSetGen(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDedicatedHostGroups().deleteWithResponse("rgcompute", "a",
            com.azure.core.util.Context.NONE);
    }
}
