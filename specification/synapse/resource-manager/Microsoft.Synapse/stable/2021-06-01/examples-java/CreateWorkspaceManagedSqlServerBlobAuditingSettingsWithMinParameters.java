import com.azure.resourcemanager.synapse.models.BlobAuditingPolicyName;
import com.azure.resourcemanager.synapse.models.BlobAuditingPolicyState;

/** Samples for WorkspaceManagedSqlServerBlobAuditingPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateWorkspaceManagedSqlServerBlobAuditingSettingsWithMinParameters.json
     */
    /**
     * Sample code: Create or update blob auditing policy of workspace managed Sql Server with minimal parameters.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createOrUpdateBlobAuditingPolicyOfWorkspaceManagedSqlServerWithMinimalParameters(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .workspaceManagedSqlServerBlobAuditingPolicies()
            .define(BlobAuditingPolicyName.DEFAULT)
            .withExistingWorkspace("wsg-7398", "testWorkspace")
            .withState(BlobAuditingPolicyState.ENABLED)
            .withStorageEndpoint("https://mystorage.blob.core.windows.net")
            .withStorageAccountAccessKey(
                "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD==")
            .create();
    }
}
