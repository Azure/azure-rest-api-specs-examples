import com.azure.core.util.Context;

/** Samples for Workflow GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/developerhub/resource-manager/Microsoft.DevHub/preview/2022-04-01-preview/examples/Workflow_Get.json
     */
    /**
     * Sample code: Get Workflow.
     *
     * @param manager Entry point to DevHubManager.
     */
    public static void getWorkflow(com.azure.resourcemanager.devhub.DevHubManager manager) {
        manager.workflows().getByResourceGroupWithResponse("resourceGroup1", "workflow1", Context.NONE);
    }
}
