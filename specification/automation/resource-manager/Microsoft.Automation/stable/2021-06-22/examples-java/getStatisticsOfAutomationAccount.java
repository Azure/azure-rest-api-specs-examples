import com.azure.core.util.Context;

/** Samples for StatisticsOperation ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/getStatisticsOfAutomationAccount.json
     */
    /**
     * Sample code: Get statistics of an automation account.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getStatisticsOfAnAutomationAccount(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.statisticsOperations().listByAutomationAccount("rg", "myAutomationAccount11", null, Context.NONE);
    }
}
