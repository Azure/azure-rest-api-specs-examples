import com.azure.core.util.Context;

/** Samples for Job ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/job/listJobsByAutomationAccount.json
     */
    /**
     * Sample code: List jobs by automation account.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listJobsByAutomationAccount(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.jobs().listByAutomationAccount("mygroup", "ContoseAutomationAccount", null, null, Context.NONE);
    }
}
