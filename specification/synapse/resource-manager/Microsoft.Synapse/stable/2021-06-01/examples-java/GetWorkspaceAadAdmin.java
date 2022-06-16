import com.azure.core.util.Context;

/** Samples for WorkspaceSqlAadAdmins Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetWorkspaceAadAdmin.json
     */
    /**
     * Sample code: Get workspace active directory admin.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getWorkspaceActiveDirectoryAdmin(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.workspaceSqlAadAdmins().getWithResponse("resourceGroup1", "workspace1", Context.NONE);
    }
}
