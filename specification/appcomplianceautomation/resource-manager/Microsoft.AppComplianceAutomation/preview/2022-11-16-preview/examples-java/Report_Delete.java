import com.azure.core.util.Context;

/** Samples for ReportOperation Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/preview/2022-11-16-preview/examples/Report_Delete.json
     */
    /**
     * Sample code: Report_Delete.
     *
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void reportDelete(
        com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.reportOperations().delete("testReportName", Context.NONE);
    }
}
