import com.azure.core.util.Context;

/** Samples for Runbook ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2018-06-30/examples/listRunbooksByAutomationAccount.json
     */
    /**
     * Sample code: List runbooks by automation account.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listRunbooksByAutomationAccount(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.runbooks().listByAutomationAccount("rg", "ContoseAutomationAccount", Context.NONE);
    }
}
