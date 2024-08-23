
/**
 * Samples for Servers ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ServerListByResourceGroupWithExpandEqualsAdministrators.json
     */
    /**
     * Sample code: List servers by resource group with $expand=administrators/activedirectory.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listServersByResourceGroupWithExpandAdministratorsActivedirectory(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServers().listByResourceGroup("sqlcrudtest-7398", null,
            com.azure.core.util.Context.NONE);
    }
}
