import com.azure.core.util.Context;

/** Samples for Servers ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2019-06-01-preview/examples/ServerListByResourceGroup.json
     */
    /**
     * Sample code: List servers by resource group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listServersByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServers().listByResourceGroup("sqlcrudtest-7398", Context.NONE);
    }
}
