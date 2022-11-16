import com.azure.core.util.Context;

/** Samples for SnapshotOperation Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/preview/2022-11-16-preview/examples/Snapshot_Get.json
     */
    /**
     * Sample code: Snapshot_Get.
     *
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void snapshotGet(
        com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.snapshotOperations().getWithResponse("testReportName", "testSnapshot", Context.NONE);
    }
}
