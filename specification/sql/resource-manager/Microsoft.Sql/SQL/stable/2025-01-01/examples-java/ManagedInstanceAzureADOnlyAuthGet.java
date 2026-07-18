
import com.azure.resourcemanager.sql.models.AuthenticationName;

/**
 * Samples for ManagedInstanceAzureADOnlyAuthentications Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceAzureADOnlyAuthGet.json
     */
    /**
     * Sample code: Gets a Azure Active Directory only authentication property.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsAAzureActiveDirectoryOnlyAuthenticationProperty(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstanceAzureADOnlyAuthentications().getWithResponse(
            "Default-SQL-SouthEastAsia", "managedInstance", AuthenticationName.DEFAULT,
            com.azure.core.util.Context.NONE);
    }
}
