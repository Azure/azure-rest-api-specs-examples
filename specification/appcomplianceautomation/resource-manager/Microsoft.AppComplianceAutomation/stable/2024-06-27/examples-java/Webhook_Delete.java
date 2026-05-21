
/**
 * Samples for Webhook Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/Webhook_Delete.json
     */
    /**
     * Sample code: Webhook_Delete.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void
        webhookDelete(com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.webhooks().deleteByResourceGroupWithResponse("testReportName", "testWebhookName",
            com.azure.core.util.Context.NONE);
    }
}
