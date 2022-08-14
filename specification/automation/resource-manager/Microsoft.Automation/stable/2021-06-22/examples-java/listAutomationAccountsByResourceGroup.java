import com.azure.core.util.Context;

/** Samples for AutomationAccount ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/listAutomationAccountsByResourceGroup.json
     */
    /**
     * Sample code: List automation accounts by resource group.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listAutomationAccountsByResourceGroup(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.automationAccounts().listByResourceGroup("rg", Context.NONE);
    }
}
