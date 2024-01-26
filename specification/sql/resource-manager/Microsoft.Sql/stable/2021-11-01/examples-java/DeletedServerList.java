
import com.azure.core.util.Context;

/** Samples for DeletedServers ListByLocation. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DeletedServerList.json
     */
    /**
     * Sample code: List deleted servers.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDeletedServers(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDeletedServers().listByLocation("japaneast", Context.NONE);
    }
}
