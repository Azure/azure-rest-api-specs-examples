
/**
 * Samples for ManagedDatabases ListByInstance.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedDatabaseListByManagedInstance.
     * json
     */
    /**
     * Sample code: List databases by managed instances.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDatabasesByManagedInstances(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabases().listByInstance("Test1", "managedInstance",
            com.azure.core.util.Context.NONE);
    }
}
