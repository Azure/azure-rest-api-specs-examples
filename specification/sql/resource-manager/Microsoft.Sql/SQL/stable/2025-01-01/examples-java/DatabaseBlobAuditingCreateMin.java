
import com.azure.resourcemanager.sql.fluent.models.DatabaseBlobAuditingPolicyInner;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyName;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyState;

/**
 * Samples for DatabaseBlobAuditingPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseBlobAuditingCreateMin.json
     */
    /**
     * Sample code: Create or update a database's blob auditing policy with minimal parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createOrUpdateADatabaseSBlobAuditingPolicyWithMinimalParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseBlobAuditingPolicies().createOrUpdateWithResponse("blobauditingtest-4799",
            "blobauditingtest-6440", "testdb", BlobAuditingPolicyName.DEFAULT,
            new DatabaseBlobAuditingPolicyInner().withState(BlobAuditingPolicyState.ENABLED)
                .withStorageEndpoint("https://mystorage.blob.core.windows.net")
                .withStorageAccountAccessKey("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }
}
