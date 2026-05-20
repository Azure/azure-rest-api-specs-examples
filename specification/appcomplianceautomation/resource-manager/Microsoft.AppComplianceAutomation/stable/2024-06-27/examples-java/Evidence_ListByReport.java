
/**
 * Samples for Evidence ListByReport.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/Evidence_ListByReport.json
     */
    /**
     * Sample code: Evidence_ListByReport.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void
        evidenceListByReport(com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.evidences().listByReport("reportName", null, null, null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
