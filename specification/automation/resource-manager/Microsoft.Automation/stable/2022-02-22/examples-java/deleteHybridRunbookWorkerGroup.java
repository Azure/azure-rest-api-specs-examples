import com.azure.core.util.Context;

/** Samples for HybridRunbookWorkerGroup Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2022-02-22/examples/deleteHybridRunbookWorkerGroup.json
     */
    /**
     * Sample code: Delete a hybrid worker group.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void deleteAHybridWorkerGroup(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.hybridRunbookWorkerGroups().deleteWithResponse("rg", "myAutomationAccount20", "myGroup", Context.NONE);
    }
}
