import com.azure.resourcemanager.synapse.models.BlobAuditingPolicyName;
import com.azure.resourcemanager.synapse.models.BlobAuditingPolicyState;
import java.util.Arrays;
import java.util.UUID;

/** Samples for WorkspaceManagedSqlServerExtendedBlobAuditingPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateWorkspaceManagedSqlServerExtendedBlobAuditingSettingsWithAllParameters.json
     */
    /**
     * Sample code: Create or update workspace managed sql server's extended blob auditing policy of with all
     * parameters.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createOrUpdateWorkspaceManagedSqlServerSExtendedBlobAuditingPolicyOfWithAllParameters(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .workspaceManagedSqlServerExtendedBlobAuditingPolicies()
            .define(BlobAuditingPolicyName.DEFAULT)
            .withExistingWorkspace("wsg-7398", "testWorkspace")
            .withPredicateExpression("object_name = 'SensitiveData'")
            .withState(BlobAuditingPolicyState.ENABLED)
            .withStorageEndpoint("https://mystorage.blob.core.windows.net")
            .withStorageAccountAccessKey(
                "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD==")
            .withRetentionDays(6)
            .withAuditActionsAndGroups(
                Arrays
                    .asList(
                        "SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP",
                        "FAILED_DATABASE_AUTHENTICATION_GROUP",
                        "BATCH_COMPLETED_GROUP"))
            .withStorageAccountSubscriptionId(UUID.fromString("00000000-1234-0000-5678-000000000000"))
            .withIsStorageSecondaryKeyInUse(false)
            .withIsAzureMonitorTargetEnabled(true)
            .create();
    }
}
