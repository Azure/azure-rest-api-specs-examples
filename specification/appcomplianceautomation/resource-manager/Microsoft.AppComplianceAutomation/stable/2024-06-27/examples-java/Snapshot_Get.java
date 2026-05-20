
/**
 * Samples for Snapshot Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/Snapshot_Get.json
     */
    /**
     * Sample code: Snapshot_Get.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void
        snapshotGet(com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.snapshots().getWithResponse("testReportName", "testSnapshot", com.azure.core.util.Context.NONE);
    }
}
