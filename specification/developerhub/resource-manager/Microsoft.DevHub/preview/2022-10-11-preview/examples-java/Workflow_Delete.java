/** Samples for Workflow Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/developerhub/resource-manager/Microsoft.DevHub/preview/2022-10-11-preview/examples/Workflow_Delete.json
     */
    /**
     * Sample code: Delete Workflow.
     *
     * @param manager Entry point to DevHubManager.
     */
    public static void deleteWorkflow(com.azure.resourcemanager.devhub.DevHubManager manager) {
        manager
            .workflows()
            .deleteByResourceGroupWithResponse("resourceGroup1", "workflow1", com.azure.core.util.Context.NONE);
    }
}
