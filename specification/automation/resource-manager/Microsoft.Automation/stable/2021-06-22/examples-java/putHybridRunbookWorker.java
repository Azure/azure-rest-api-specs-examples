/** Samples for HybridRunbookWorkers Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/putHybridRunbookWorker.json
     */
    /**
     * Sample code: Create a V2 hybrid runbook worker.
     *
     * @param manager Entry point to AutomationManager.
     */
    public static void createAV2HybridRunbookWorker(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager
            .hybridRunbookWorkers()
            .define("c010ad12-ef14-4a2a-aa9e-ef22c4745ddd")
            .withExistingHybridRunbookWorkerGroup("rg", "testaccount", "TestHybridGroup")
            .withVmResourceId(
                "/subscriptions/vmsubid/resourceGroups/vmrg/providers/Microsoft.Compute/virtualMachines/vmname")
            .create();
    }
}
