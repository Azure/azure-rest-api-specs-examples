
/**
 * Samples for DatabaseBlobAuditingPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DatabaseBlobAuditingGet.json
     */
    /**
     * Sample code: Get a database's blob auditing policy.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getADatabaseSBlobAuditingPolicy(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseBlobAuditingPolicies().getWithResponse(
            "blobauditingtest-6852", "blobauditingtest-2080", "testdb", com.azure.core.util.Context.NONE);
    }
}
