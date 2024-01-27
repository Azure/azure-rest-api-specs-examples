
import com.azure.core.util.Context;

/** Samples for ManagedDatabaseSecurityEvents ListByDatabase. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedDatabaseSecurityEventsGetMin.
     * json
     */
    /**
     * Sample code: Get the managed database's security events with minimal parameters.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getTheManagedDatabaseSSecurityEventsWithMinimalParameters(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabaseSecurityEvents().listByDatabase("testrg",
            "testcl", "database1", null, null, null, null, Context.NONE);
    }
}
