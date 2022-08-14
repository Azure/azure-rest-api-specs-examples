import com.azure.core.util.Context;

/** Samples for Usages ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/getUsagesOfAutomationAccount.json
     */
    /**
     * Sample code: Get usages of an automation account.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getUsagesOfAnAutomationAccount(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.usages().listByAutomationAccount("rg", "myAutomationAccount11", Context.NONE);
    }
}
