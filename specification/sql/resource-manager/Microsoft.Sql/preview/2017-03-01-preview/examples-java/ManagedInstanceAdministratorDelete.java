import com.azure.core.util.Context;

/** Samples for ManagedInstanceAdministrators Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/ManagedInstanceAdministratorDelete.json
     */
    /**
     * Sample code: Delete administrator of managed instance.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAdministratorOfManagedInstance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getManagedInstanceAdministrators()
            .delete("Default-SQL-SouthEastAsia", "managedInstance", Context.NONE);
    }
}
