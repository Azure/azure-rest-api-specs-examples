
/**
 * Samples for ManagedInstanceAdministrators ListByInstance.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ManagedInstanceAdministratorListByInstance.json
     */
    /**
     * Sample code: List administrators of managed instance.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAdministratorsOfManagedInstance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstanceAdministrators()
            .listByInstance("Default-SQL-SouthEastAsia", "managedInstance", com.azure.core.util.Context.NONE);
    }
}
