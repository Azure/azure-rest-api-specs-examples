
import com.azure.core.util.Context;

/** Samples for Servers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ServerDelete.json
     */
    /**
     * Sample code: Delete server.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteServer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServers().delete("sqlcrudtest-7398", "sqlcrudtest-6661",
            Context.NONE);
    }
}
