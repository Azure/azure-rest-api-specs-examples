import com.azure.core.util.Context;

/** Samples for Webhook Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2015-10-31/examples/getWebhook.json
     */
    /**
     * Sample code: Get webhook.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getWebhook(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.webhooks().getWithResponse("rg", "myAutomationAccount33", "TestWebhook", Context.NONE);
    }
}
