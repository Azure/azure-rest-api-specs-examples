import com.azure.core.util.Context;

/** Samples for Servers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2019-06-01-preview/examples/ServerList.json
     */
    /**
     * Sample code: List servers.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listServers(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServers().list(Context.NONE);
    }
}
