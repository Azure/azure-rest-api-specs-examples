import com.azure.core.util.Context;

/** Samples for DatabaseBlobAuditingPolicies ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/DatabaseAuditingSettingsList.json
     */
    /**
     * Sample code: List audit settings of a database.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAuditSettingsOfADatabase(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDatabaseBlobAuditingPolicies()
            .listByDatabase("blobauditingtest-6852", "blobauditingtest-2080", "testdb", Context.NONE);
    }
}
