
/**
 * Samples for DrillResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/DrillResources_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: DrillResources_Get_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void drillResourcesGetMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.drillResources().getWithResponse("sampleServiceGroupName", "drill1",
            "b6378181-9dc0-4a43-8e09-97a8b08aabaa", com.azure.core.util.Context.NONE);
    }
}
