
import com.azure.core.util.Context;

/** Samples for ServerKeys ListByServer. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerKeyList.json
     */
    /**
     * Sample code: List the server keys by server.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listTheServerKeysByServer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerKeys().listByServer("sqlcrudtest-7398",
            "sqlcrudtest-4645", Context.NONE);
    }
}
