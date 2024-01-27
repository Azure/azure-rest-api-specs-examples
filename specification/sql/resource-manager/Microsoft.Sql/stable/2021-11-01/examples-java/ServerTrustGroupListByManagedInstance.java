
import com.azure.core.util.Context;

/** Samples for ServerTrustGroups ListByInstance. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerTrustGroupListByManagedInstance
     * .json
     */
    /**
     * Sample code: List server trust groups by managed instance.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listServerTrustGroupsByManagedInstance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerTrustGroups().listByInstance("Default-SQL-SouthEastAsia",
            "managedInstance-1", Context.NONE);
    }
}
