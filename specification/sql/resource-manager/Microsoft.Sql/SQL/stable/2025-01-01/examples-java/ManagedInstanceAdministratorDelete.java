
import com.azure.resourcemanager.sql.models.AdministratorName;

/**
 * Samples for ManagedInstanceAdministrators Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceAdministratorDelete.json
     */
    /**
     * Sample code: Delete administrator of managed instance.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteAdministratorOfManagedInstance(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceAdministrators().delete("Default-SQL-SouthEastAsia",
            "managedInstance", AdministratorName.ACTIVE_DIRECTORY, com.azure.core.util.Context.NONE);
    }
}
