
/**
 * Samples for DrillResources List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/DrillResources_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: DrillResources_List_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void drillResourcesListMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.drillResources().list("sampleServiceGroupName", "drill1", "xntbyoswztnmvitj", 69,
            com.azure.core.util.Context.NONE);
    }
}
