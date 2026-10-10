
/**
 * Samples for DrillRunResources List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/DrillRunResources_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: DrillRunResources_List_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void drillRunResourcesListMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.drillRunResources().list("sampleServiceGroupName", "drill1", "ca92602e-53bf-43d2-ae62-d3fc940474b3",
            com.azure.core.util.Context.NONE);
    }
}
