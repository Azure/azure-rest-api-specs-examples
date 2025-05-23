
import com.azure.resourcemanager.automation.models.HybridRunbookWorkerMoveParameters;

/**
 * Samples for HybridRunbookWorkers Move.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/automation/resource-manager/Microsoft.Automation/stable/2021-06-22/examples/moveHybridRunbookWorker
     * .json
     */
    /**
     * Sample code: Move a V2 hybrid worker to a different group.
     * 
     * @param manager Entry point to AutomationManager.
     */
    public static void
        moveAV2HybridWorkerToADifferentGroup(com.azure.resourcemanager.automation.AutomationManager manager) {
        manager.hybridRunbookWorkers().moveWithResponse("rg", "testaccount", "TestHybridGroup",
            "c010ad12-ef14-4a2a-aa9e-ef22c4745ddd",
            new HybridRunbookWorkerMoveParameters().withHybridRunbookWorkerGroupName("TestHybridGroup2"),
            com.azure.core.util.Context.NONE);
    }
}
