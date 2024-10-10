
/**
 * Samples for AutomationAccount Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/deleteAutomationAccount
     * .json
     */
    /**
     * Sample code: Delete automation account.
     * 
     * @param manager Entry point to AutomationManager.
     */
    public static void deleteAutomationAccount(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.automationAccounts().deleteByResourceGroupWithResponse("rg", "myAutomationAccount9",
            com.azure.core.util.Context.NONE);
    }
}
