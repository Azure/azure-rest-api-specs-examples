
/**
 * Samples for RecoverableManagedDatabases ListByInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ListRecoverableManagedDatabasesByServer.json
     */
    /**
     * Sample code: List recoverable databases by managed instances.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        listRecoverableDatabasesByManagedInstances(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getRecoverableManagedDatabases().listByInstance("Test1",
            "managedInstance", com.azure.core.util.Context.NONE);
    }
}
