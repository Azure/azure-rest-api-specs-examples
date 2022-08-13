import com.azure.core.util.Context;

/** Samples for HybridRunbookWorkerGroup Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2022-02-22/examples/getHybridRunbookWorkerGroup.json
     */
    /**
     * Sample code: Get a hybrid worker group.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void getAHybridWorkerGroup(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.hybridRunbookWorkerGroups().getWithResponse("rg", "testaccount", "TestHybridGroup", Context.NONE);
    }
}
