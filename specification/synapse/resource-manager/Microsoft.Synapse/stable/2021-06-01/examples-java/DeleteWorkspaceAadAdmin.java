/** Samples for WorkspaceSqlAadAdmins Delete. */
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
        manager.workspaceSqlAadAdmins().delete("resourceGroup1", "workspace1", com.azure.core.util.Context.NONE);
    }
}
