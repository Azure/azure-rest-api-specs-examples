
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyName;

/**
 * Samples for ExtendedDatabaseBlobAuditingPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ExtendedDatabaseBlobAuditingGet.json
     */
    /**
     * Sample code: Get an extended database's blob auditing policy.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getAnExtendedDatabaseSBlobAuditingPolicy(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getExtendedDatabaseBlobAuditingPolicies().getWithResponse("blobauditingtest-6852",
            "blobauditingtest-2080", "testdb", BlobAuditingPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
