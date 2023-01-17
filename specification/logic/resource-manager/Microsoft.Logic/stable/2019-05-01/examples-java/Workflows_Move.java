import com.azure.resourcemanager.logic.models.WorkflowReference;

/** Samples for Workflows Move. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/Workflows_Move.json
     */
    /**
     * Sample code: Move a workflow.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void moveAWorkflow(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .workflows()
            .move(
                "testResourceGroup",
                "testWorkflow",
                new WorkflowReference()
                    .withId(
                        "subscriptions/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/newResourceGroup/providers/Microsoft.Logic/workflows/newWorkflowName"),
                com.azure.core.util.Context.NONE);
    }
}
