
import com.azure.core.util.Context;

/** Samples for ServerTrustGroups ListByLocation. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerTrustGroupList.json
     */
    /**
     * Sample code: List server trust groups.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listServerTrustGroups(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerTrustGroups().listByLocation("Default", "Japan East",
            Context.NONE);
    }
}
