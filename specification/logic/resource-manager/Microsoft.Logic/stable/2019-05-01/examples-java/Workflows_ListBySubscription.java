/** Samples for Workflows List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_ListBySubscription.json
     */
    /**
     * Sample code: List all workflows in a subscription.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void listAllWorkflowsInASubscription(com.azure.resourcemanager.logic.LogicManager manager) {
        manager.workflows().list(null, null, com.azure.core.util.Context.NONE);
    }
}
