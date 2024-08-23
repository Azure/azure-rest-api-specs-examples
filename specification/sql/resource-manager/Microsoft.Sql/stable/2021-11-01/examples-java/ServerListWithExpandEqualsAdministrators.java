
/**
 * Samples for Servers List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ServerListWithExpandEqualsAdministrators.json
     */
    /**
     * Sample code: List servers with $expand=administrators/activedirectory.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listServersWithExpandAdministratorsActivedirectory(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServers().list(null, com.azure.core.util.Context.NONE);
    }
}
