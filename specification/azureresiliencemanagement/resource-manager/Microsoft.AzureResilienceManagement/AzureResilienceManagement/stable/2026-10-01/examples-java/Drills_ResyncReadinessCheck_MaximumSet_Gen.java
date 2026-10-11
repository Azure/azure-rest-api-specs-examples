
/**
 * Samples for Drills ResyncReadinessCheck.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Drills_ResyncReadinessCheck_MaximumSet_Gen.json
     */
    /**
     * Sample code: Drills_ResyncReadinessCheck_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void drillsResyncReadinessCheckMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.drills().resyncReadinessCheck("sampleServiceGroupName", "qmn", "drill1",
            com.azure.core.util.Context.NONE);
    }
}
