
/**
 * Samples for HybridRunbookWorkers Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/getHybridRunbookWorker.
     * json
     */
    /**
     * Sample code: Get a V2 hybrid runbook worker.
     * 
     * @param manager Entry point to AutomationManager.
     */
    public static void getAV2HybridRunbookWorker(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.hybridRunbookWorkers().getWithResponse("rg", "testaccount", "TestHybridGroup",
            "c010ad12-ef14-4a2a-aa9e-ef22c4745ddd", com.azure.core.util.Context.NONE);
    }
}
