
import com.azure.resourcemanager.sql.fluent.models.ExtendedDatabaseBlobAuditingPolicyInner;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyName;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyState;

/**
 * Samples for ExtendedDatabaseBlobAuditingPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ExtendedDatabaseBlobAuditingCreateMin.json
     */
    /**
     * Sample code: Create or update an extended database's blob auditing policy with minimal parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createOrUpdateAnExtendedDatabaseSBlobAuditingPolicyWithMinimalParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getExtendedDatabaseBlobAuditingPolicies().createOrUpdateWithResponse(
            "blobauditingtest-4799", "blobauditingtest-6440", "testdb", BlobAuditingPolicyName.DEFAULT,
            new ExtendedDatabaseBlobAuditingPolicyInner().withState(BlobAuditingPolicyState.ENABLED)
                .withStorageEndpoint("https://mystorage.blob.core.windows.net")
                .withStorageAccountAccessKey("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }
}
