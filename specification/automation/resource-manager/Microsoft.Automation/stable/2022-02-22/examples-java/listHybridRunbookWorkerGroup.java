import com.azure.core.util.Context;

/** Samples for HybridRunbookWorkerGroup ListByAutomationAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2022-02-22/examples/listHybridRunbookWorkerGroup.json
     */
    /**
     * Sample code: List hybrid worker groups by Automation Account.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listHybridWorkerGroupsByAutomationAccount(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.hybridRunbookWorkerGroups().listByAutomationAccount("rg", "testaccount", null, Context.NONE);
    }
}
