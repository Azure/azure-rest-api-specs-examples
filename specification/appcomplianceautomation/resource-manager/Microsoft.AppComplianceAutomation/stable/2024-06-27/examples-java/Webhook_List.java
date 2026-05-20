
/**
 * Samples for Webhook List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/Webhook_List.json
     */
    /**
     * Sample code: Webhook_List.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void
        webhookList(com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.webhooks().list("testReportName", "1", 100, null, null, null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
