import com.azure.core.util.Context;

/** Samples for WorkspaceAadAdmins Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/DeleteWorkspaceAadAdmin.json
     */
    /**
     * Sample code: Delete workspace active directory admin.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void deleteWorkspaceActiveDirectoryAdmin(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.workspaceAadAdmins().delete("resourceGroup1", "workspace1", Context.NONE);
    }
}
