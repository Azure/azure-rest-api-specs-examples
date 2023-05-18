/** Samples for Workflow ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/developerhub/resource-manager/Microsoft.DevHub/preview/2022-10-11-preview/examples/Workflow_ListByResourceGroup.json
     */
    /**
     * Sample code: List Workflows.
     *
     * @param manager Entry point to DevHubManager.
     */
    public static void listWorkflows(com.azure.resourcemanager.devhub.DevHubManager manager) {
        manager
            .workflows()
            .listByResourceGroup(
                "resourceGroup1",
                "/subscriptions/subscriptionId1/resourcegroups/resourceGroup1/providers/Microsoft.ContainerService/managedClusters/cluster1",
                com.azure.core.util.Context.NONE);
    }
}
