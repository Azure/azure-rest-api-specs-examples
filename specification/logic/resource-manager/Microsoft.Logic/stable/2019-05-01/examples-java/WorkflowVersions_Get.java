/** Samples for WorkflowVersions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/WorkflowVersions_Get.json
     */
    /**
     * Sample code: Get a workflow version.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getAWorkflowVersion(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflowVersions()
            .getWithResponse(
                "test-resource-group", "test-workflow", "08586676824806722526", com.azure.core.util.Context.NONE);
    }
}
