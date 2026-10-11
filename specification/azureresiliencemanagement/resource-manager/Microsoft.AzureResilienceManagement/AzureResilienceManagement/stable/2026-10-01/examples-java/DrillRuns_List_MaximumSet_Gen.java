
/**
 * Samples for DrillRuns List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/DrillRuns_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: DrillRuns_List_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void
        drillRunsListMaximumSet(com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.drillRuns().list("sampleServiceGroupName", "drill1", com.azure.core.util.Context.NONE);
    }
}
