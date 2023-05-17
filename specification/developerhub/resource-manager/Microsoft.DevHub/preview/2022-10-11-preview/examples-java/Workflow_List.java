/** Samples for Workflow List. */
public final class Main {
    /*
     * x-ms-original-file: specification/developerhub/resource-manager/Microsoft.DevHub/preview/2022-10-11-preview/examples/Workflow_List.json
     */
    /**
     * Sample code: List Workflows.
     *
     * @param manager Entry point to DevHubManager.
     */
    public static void listWorkflows(com.azure.resourcemanager.devhub.DevHubManager manager) {
        manager.workflows().list(com.azure.core.util.Context.NONE);
    }
}
