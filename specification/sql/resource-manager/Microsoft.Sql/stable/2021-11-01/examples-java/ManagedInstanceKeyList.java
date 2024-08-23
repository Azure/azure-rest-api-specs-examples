
/**
 * Samples for ManagedInstanceKeys ListByInstance.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedInstanceKeyList.json
     */
    /**
     * Sample code: List the keys for a managed instance.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listTheKeysForAManagedInstance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstanceKeys().listByInstance("sqlcrudtest-7398",
            "sqlcrudtest-4645", null, com.azure.core.util.Context.NONE);
    }
}
