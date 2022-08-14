import com.azure.core.util.Context;

/** Samples for Webhook GenerateUri. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2015-10-31/examples/webhookGenerateUri.json
     */
    /**
     * Sample code: Generate webhook uri.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void generateWebhookUri(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.webhooks().generateUriWithResponse("rg", "myAutomationAccount33", Context.NONE);
    }
}
