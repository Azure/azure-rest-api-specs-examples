
/**
 * Samples for Evidence Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/Evidence_Delete.json
     */
    /**
     * Sample code: Evidence_Delete.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void
        evidenceDelete(com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.evidences().deleteByResourceGroupWithResponse("testReportName", "evidence1",
            com.azure.core.util.Context.NONE);
    }
}
