
import com.azure.resourcemanager.synapse.models.BlobAuditingPolicyState;
import java.util.Arrays;
import java.util.UUID;

/**
 * Samples for SqlPoolBlobAuditingPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/
     * CreateOrUpdateSqlPoolBlobAuditingWithAllParameters.json
     */
    /**
     * Sample code: Create or update a database's blob auditing policy with all parameters.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void createOrUpdateADatabaseSBlobAuditingPolicyWithAllParameters(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.sqlPoolBlobAuditingPolicies().define()
            .withExistingSqlPool("blobauditingtest-4799", "blobauditingtest-6440", "testdb")
            .withState(BlobAuditingPolicyState.ENABLED).withStorageEndpoint("https://mystorage.blob.core.windows.net")
            .withStorageAccountAccessKey(
                "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD==")
            .withRetentionDays(6)
            .withAuditActionsAndGroups(Arrays.asList("DATABASE_LOGOUT_GROUP", "DATABASE_ROLE_MEMBER_CHANGE_GROUP",
                "UPDATE on database::TestDatabaseName by public"))
            .withStorageAccountSubscriptionId(UUID.fromString("00000000-1234-0000-5678-000000000000"))
            .withIsStorageSecondaryKeyInUse(false).withIsAzureMonitorTargetEnabled(true).create();
    }
}
