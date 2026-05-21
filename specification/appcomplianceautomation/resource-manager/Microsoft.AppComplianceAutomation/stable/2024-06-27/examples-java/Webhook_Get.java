
/**
 * Samples for Webhook Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-06-27/Webhook_Get.json
     */
    /**
     * Sample code: Webhook_Get.
     * 
     * @param manager Entry point to AppComplianceAutomationManager.
     */
    public static void
        webhookGet(com.azure.resourcemanager.appcomplianceautomation.AppComplianceAutomationManager manager) {
        manager.webhooks().getWithResponse("testReportName", "testWebhookName", com.azure.core.util.Context.NONE);
    }
}
