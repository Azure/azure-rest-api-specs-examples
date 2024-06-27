
/**
 * Samples for Webhook CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appcomplianceautomation/resource-manager/Microsoft.AppComplianceAutomation/stable/2024-06-27/
     * examples/Webhook_CreateOrUpdate.json
     */
    /**
     * Sample code: Webhook_CreateOrUpdate.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void webhookCreateOrUpdate(
        com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.webhooks().createOrUpdateWithResponse("testReportName", "testWebhookName", null,
            com.azure.core.util.Context.NONE);
    }
}
