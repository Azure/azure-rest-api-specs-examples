import com.azure.core.util.Context;

/** Samples for Credential ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listCredentialsByAutomationAccount.json
     */
    /**
     * Sample code: List credentials by automation account, first 100.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listCredentialsByAutomationAccountFirst100(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.credentials().listByAutomationAccount("rg", "myAutomationAccount20", Context.NONE);
    }
}
