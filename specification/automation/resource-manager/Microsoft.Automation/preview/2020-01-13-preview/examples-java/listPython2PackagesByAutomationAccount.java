import com.azure.core.util.Context;

/** Samples for Python2Package ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listPython2PackagesByAutomationAccount.json
     */
    /**
     * Sample code: List python 2 packages by automation account.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listPython2PackagesByAutomationAccount(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.python2Packages().listByAutomationAccount("rg", "myAutomationAccount33", Context.NONE);
    }
}
