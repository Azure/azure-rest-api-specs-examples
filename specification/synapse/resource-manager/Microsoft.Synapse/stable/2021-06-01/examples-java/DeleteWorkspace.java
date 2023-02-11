/** Samples for Workspaces Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeleteWorkspace.json
     */
    /**
     * Sample code: Delete a workspace.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void deleteAWorkspace(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.workspaces().delete("resourceGroup1", "workspace1", com.azure.core.util.Context.NONE);
    }
}
