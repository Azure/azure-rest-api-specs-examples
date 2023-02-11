/** Samples for Workspaces GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetWorkspace.json
     */
    /**
     * Sample code: Get a workspace.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getAWorkspace(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .workspaces()
            .getByResourceGroupWithResponse("resourceGroup1", "workspace1", com.azure.core.util.Context.NONE);
    }
}
