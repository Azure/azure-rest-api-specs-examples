/** Samples for WorkflowVersions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowVersions_List.json
     */
    /**
     * Sample code: List a workflows versions.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void listAWorkflowsVersions(com.azure.resourcemanager.logic.LogicManager manager) {
        manager.workflowVersions().list("test-resource-group", "test-workflow", null, com.azure.core.util.Context.NONE);
    }
}
