Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.4/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.synapse.models.BlobAuditingPolicyName;
import com.azure.resourcemanager.synapse.models.BlobAuditingPolicyState;
import java.util.Arrays;
import java.util.UUID;

/** Samples for WorkspaceManagedSqlServerBlobAuditingPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateWorkspaceManagedSqlServerBlobAuditingSettingsWithAllParameters.json
     */
    /**
     * Sample code: Create or update blob auditing policy of workspace SQL Server with all parameters.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createOrUpdateBlobAuditingPolicyOfWorkspaceSQLServerWithAllParameters(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .workspaceManagedSqlServerBlobAuditingPolicies()
            .define(BlobAuditingPolicyName.DEFAULT)
            .withExistingWorkspace("wsg-7398", "testWorkspace")
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
            .withQueueDelayMs(4000)
            .create();
    }
}
```
