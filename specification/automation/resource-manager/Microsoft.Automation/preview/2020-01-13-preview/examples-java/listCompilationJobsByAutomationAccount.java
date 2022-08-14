import com.azure.core.util.Context;

/** Samples for DscCompilationJob ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/preview/2020-01-13-preview/examples/listCompilationJobsByAutomationAccount.json
     */
    /**
     * Sample code: List DSC Compilation job in Automation Account.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listDSCCompilationJobInAutomationAccount(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.dscCompilationJobs().listByAutomationAccount("rg", "myAutomationAccount33", null, Context.NONE);
    }
}
