
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.AdministratorName;

/** Samples for ManagedInstanceAdministrators Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedInstanceAdministratorGet.json
     */
    /**
     * Sample code: Get administrator of managed instance.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getAdministratorOfManagedInstance(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstanceAdministrators().getWithResponse(
            "Default-SQL-SouthEastAsia", "managedInstance", AdministratorName.ACTIVE_DIRECTORY, Context.NONE);
    }
}
