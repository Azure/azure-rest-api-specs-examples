import com.azure.core.util.Context;

/** Samples for HybridRunbookWorkers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/deleteHybridRunbookWorker.json
     */
    /**
     * Sample code: Delete a V2 hybrid runbook worker.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void deleteAV2HybridRunbookWorker(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .hybridRunbookWorkers()
            .deleteWithResponse(
                "rg", "myAutomationAccount20", "myGroup", "c010ad12-ef14-4a2a-aa9e-ef22c4745ddd", Context.NONE);
    }
}
