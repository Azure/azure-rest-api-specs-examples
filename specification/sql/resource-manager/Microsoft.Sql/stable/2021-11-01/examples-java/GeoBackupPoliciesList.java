
import com.azure.core.util.Context;

/** Samples for GeoBackupPolicies List. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/GeoBackupPoliciesList.json
     */
    /**
     * Sample code: List Geo backup policies for the given database resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listGeoBackupPoliciesForTheGivenDatabaseResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getGeoBackupPolicies().list("sqlcrudtest-4799", "sqlcrudtest-5961",
            "testdw", Context.NONE);
    }
}
