
/**
 * Samples for WorkflowRunActionRepetitionsRequestHistories List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/
     * WorkflowRunActionRepetitionsRequestHistories_List.json
     */
    /**
     * Sample code: List repetition request history.
     * 
     * @param manager Entry point to LogicManager.
     */
    public static void listRepetitionRequestHistory(com.azure.resourcemanager.logic.LogicManager manager) {
        manager.workflowRunActionRepetitionsRequestHistories().list("test-resource-group", "test-workflow",
            "08586776228332053161046300351", "HTTP_Webhook", "000001", com.azure.core.util.Context.NONE);
    }
}
