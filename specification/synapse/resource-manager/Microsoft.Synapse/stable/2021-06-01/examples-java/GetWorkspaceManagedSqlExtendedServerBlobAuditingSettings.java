import com.azure.resourcemanager.synapse.models.BlobAuditingPolicyName;

/** Samples for WorkspaceManagedSqlServerExtendedBlobAuditingPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetWorkspaceManagedSqlExtendedServerBlobAuditingSettings.json
     */
    /**
     * Sample code: Get workspace managed sql servers' extended blob auditing settings.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getWorkspaceManagedSqlServersExtendedBlobAuditingSettings(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .workspaceManagedSqlServerExtendedBlobAuditingPolicies()
            .getWithResponse(
                "wsg-7398", "testWorkspace", BlobAuditingPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
