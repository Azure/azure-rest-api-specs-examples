import com.azure.core.util.Context;

/** Samples for Webhook Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2015-10-31/examples/deleteWebhook.json
     */
    /**
     * Sample code: Delete webhook.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void deleteWebhook(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.webhooks().deleteWithResponse("rg", "myAutomationAccount33", "TestWebhook", Context.NONE);
    }
}
