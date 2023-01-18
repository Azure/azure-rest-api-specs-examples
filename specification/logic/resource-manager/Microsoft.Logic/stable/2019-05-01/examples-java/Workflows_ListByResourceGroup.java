/** Samples for Workflows ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_ListByResourceGroup.json
     */
    /**
     * Sample code: List all workflows in a resource group.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void listAllWorkflowsInAResourceGroup(com.azure.resourcemanager.logic.LogicManager manager) {
        manager.workflows().listByResourceGroup("test-resource-group", null, null, com.azure.core.util.Context.NONE);
    }
}
