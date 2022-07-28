import com.azure.core.util.Context;

/** Samples for ManagedInstanceAdministrators ListByInstance. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/ManagedInstanceAdministratorListByInstance.json
     */
    /**
     * Sample code: List administrators of managed instance.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAdministratorsOfManagedInstance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getManagedInstanceAdministrators()
            .listByInstance("Default-SQL-SouthEastAsia", "managedInstance", Context.NONE);
    }
}
