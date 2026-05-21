
/**
 * Samples for Report GetScopingQuestions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/Report_GetScopingQuestions.json
     */
    /**
     * Sample code: Report_GetScopingQuestions.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void reportGetScopingQuestions(
        com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.reports().getScopingQuestionsWithResponse("testReportName", com.azure.core.util.Context.NONE);
    }
}
