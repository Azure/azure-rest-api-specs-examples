
/**
 * Samples for Webhook Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/Webhook_Update.json
     */
    /**
     * Sample code: Webhook_Update.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void
        webhookUpdate(com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.webhooks().updateWithResponse("testReportName", "testWebhookName", null,
            com.azure.core.util.Context.NONE);
    }
}
