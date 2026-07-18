
import com.azure.resourcemanager.sql.models.AuthenticationName;

/**
 * Samples for ManagedInstanceAzureADOnlyAuthentications Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceAzureADOnlyAuthDelete.json
     */
    /**
     * Sample code: Deletes Azure Active Directory only authentication object.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        deletesAzureActiveDirectoryOnlyAuthenticationObject(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceAzureADOnlyAuthentications().delete("Default-SQL-SouthEastAsia",
            "managedInstance", AuthenticationName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
