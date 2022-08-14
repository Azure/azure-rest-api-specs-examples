import com.azure.core.util.Context;

/** Samples for Keys ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/listAutomationAccountKeys.json
     */
    /**
     * Sample code: Get lists of an automation account.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getListsOfAnAutomationAccount(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.keys().listByAutomationAccountWithResponse("rg", "MyAutomationAccount", Context.NONE);
    }
}
