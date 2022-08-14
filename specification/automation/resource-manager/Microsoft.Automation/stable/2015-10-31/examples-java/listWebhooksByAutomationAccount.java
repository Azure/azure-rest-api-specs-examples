import com.azure.core.util.Context;

/** Samples for Webhook ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2015-10-31/examples/listWebhooksByAutomationAccount.json
     */
    /**
     * Sample code: List webhooks by Automation Account.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listWebhooksByAutomationAccount(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.webhooks().listByAutomationAccount("rg", "myAutomationAccount33", null, Context.NONE);
    }
}
