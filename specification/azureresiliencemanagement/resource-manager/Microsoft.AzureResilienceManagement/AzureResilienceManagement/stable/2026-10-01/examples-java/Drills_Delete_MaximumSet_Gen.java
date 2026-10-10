
/**
 * Samples for Drills Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Drills_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Drills_Delete_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void
        drillsDeleteMaximumSet(com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.drills().delete("sampleServiceGroupName", "drill1", com.azure.core.util.Context.NONE);
    }
}
