import com.azure.core.util.Context;

/** Samples for DscNode ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listAllDscNodesByAutomationAccount.json
     */
    /**
     * Sample code: List DSC nodes by Automation Account.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listDSCNodesByAutomationAccount(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.dscNodes().listByAutomationAccount("rg", "myAutomationAccount33", null, null, null, null, Context.NONE);
    }
}
