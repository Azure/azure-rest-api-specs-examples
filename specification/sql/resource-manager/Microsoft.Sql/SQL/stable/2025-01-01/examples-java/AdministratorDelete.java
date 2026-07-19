
import com.azure.resourcemanager.sql.models.AdministratorName;

/**
 * Samples for ServerAzureADAdministrators Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/AdministratorDelete.json
     */
    /**
     * Sample code: Delete Azure Active Directory administrator.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void deleteAzureActiveDirectoryAdministrator(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerAzureADAdministrators().delete("sqlcrudtest-4799", "sqlcrudtest-6440",
            AdministratorName.ACTIVE_DIRECTORY, com.azure.core.util.Context.NONE);
    }
}
