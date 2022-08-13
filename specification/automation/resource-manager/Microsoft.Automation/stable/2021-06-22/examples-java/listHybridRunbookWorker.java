import com.azure.core.util.Context;

/** Samples for HybridRunbookWorkers ListByHybridRunbookWorkerGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/listHybridRunbookWorker.json
     */
    /**
     * Sample code: List hybrid workers by hybrid runbook worker group.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void listHybridWorkersByHybridRunbookWorkerGroup(
        com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .hybridRunbookWorkers()
            .listByHybridRunbookWorkerGroup("rg", "testaccount", "TestHybridGroup", null, Context.NONE);
    }
}
