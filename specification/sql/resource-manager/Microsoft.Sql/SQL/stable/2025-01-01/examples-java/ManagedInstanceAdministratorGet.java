
import com.azure.resourcemanager.sql.models.AdministratorName;

/**
 * Samples for ManagedInstanceAdministrators Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceAdministratorGet.json
     */
    /**
     * Sample code: Get administrator of managed instance.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAdministratorOfManagedInstance(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceAdministrators().getWithResponse("Default-SQL-SouthEastAsia",
            "managedInstance", AdministratorName.ACTIVE_DIRECTORY, com.azure.core.util.Context.NONE);
    }
}
