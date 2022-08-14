import com.azure.core.util.Context;
import com.azure.resourcemanager.automation.models.Webhook;

/** Samples for Webhook Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2015-10-31/examples/updateWebhook.json
     */
    /**
     * Sample code: Update webhook.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void updateWebhook(com.azure.resourcemanager.automation.AutomationManager manager) {
        Webhook resource =
            manager.webhooks().getWithResponse("rg", "myAutomationAccount33", "TestWebhook", Context.NONE).getValue();
        resource.update().withName("TestWebhook").withIsEnabled(false).withDescription("updated webhook").apply();
    }
}
