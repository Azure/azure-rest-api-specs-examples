import com.azure.core.util.Context;

/** Samples for GeoBackupPolicies ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2014-04-01/examples/GeoBackupPoliciesList.json
     */
    /**
     * Sample code: List geo backup policies.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listGeoBackupPolicies(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getGeoBackupPolicies()
            .listByDatabase("sqlcrudtest-4799", "sqlcrudtest-5961", "testdw", Context.NONE);
    }
}
