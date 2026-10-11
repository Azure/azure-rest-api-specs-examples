
/**
 * Samples for Drills List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Drills_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Drills_List_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void
        drillsListMaximumSet(com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.drills().list("sampleServiceGroupName", "xntbyoswztnmvitj", 69, com.azure.core.util.Context.NONE);
    }
}
