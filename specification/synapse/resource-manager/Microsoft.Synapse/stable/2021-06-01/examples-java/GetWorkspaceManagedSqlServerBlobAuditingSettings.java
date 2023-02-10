import com.azure.resourcemanager.synapse.models.BlobAuditingPolicyName;

/** Samples for WorkspaceManagedSqlServerBlobAuditingPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetWorkspaceManagedSqlServerBlobAuditingSettings.json
     */
    /**
     * Sample code: Get blob auditing setting of workspace managed sql Server.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getBlobAuditingSettingOfWorkspaceManagedSqlServer(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .workspaceManagedSqlServerBlobAuditingPolicies()
            .getWithResponse(
                "wsg-7398", "testWorkspace", BlobAuditingPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
