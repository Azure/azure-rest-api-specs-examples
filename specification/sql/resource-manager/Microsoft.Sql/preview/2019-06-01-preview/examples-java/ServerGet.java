import com.azure.core.util.Context;

/** Samples for Servers GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2019-06-01-preview/examples/ServerGet.json
     */
    /**
     * Sample code: Get server.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getServer(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getServers()
            .getByResourceGroupWithResponse("sqlcrudtest-7398", "sqlcrudtest-4645", Context.NONE);
    }
}
