
/**
 * Samples for ExtendedDatabaseBlobAuditingPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ExtendedDatabaseBlobAuditingGet.json
     */
    /**
     * Sample code: Get an extended database's blob auditing policy.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAnExtendedDatabaseSBlobAuditingPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getExtendedDatabaseBlobAuditingPolicies().getWithResponse(
            "blobauditingtest-6852", "blobauditingtest-2080", "testdb", com.azure.core.util.Context.NONE);
    }
}
