
import com.azure.resourcemanager.synapse.fluent.models.WorkspaceAadAdminInfoInner;

/**
 * Samples for WorkspaceSqlAadAdmins CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-03-01/examples/
     * CreateOrUpdateWorkspaceAadAdmin.json
     */
    /**
     * Sample code: Create or update workspace active directory admin.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void
        createOrUpdateWorkspaceActiveDirectoryAdmin(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.workspaceSqlAadAdmins().createOrUpdate("resourceGroup1", "workspace1",
            new WorkspaceAadAdminInfoInner().withTenantId("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c")
                .withLogin("bob@contoso.com").withAdministratorType("ActiveDirectory")
                .withSid("c6b82b90-a647-49cb-8a62-0d2d3cb7ac7c"),
            com.azure.core.util.Context.NONE);
    }
}
