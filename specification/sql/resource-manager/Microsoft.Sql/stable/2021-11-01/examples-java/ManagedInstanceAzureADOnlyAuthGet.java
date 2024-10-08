
import com.azure.resourcemanager.sql.models.AuthenticationName;

/**
 * Samples for ManagedInstanceAzureADOnlyAuthentications Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedInstanceAzureADOnlyAuthGet.
     * json
     */
    /**
     * Sample code: Gets a Azure Active Directory only authentication property.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getsAAzureActiveDirectoryOnlyAuthenticationProperty(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstanceAzureADOnlyAuthentications().getWithResponse(
            "Default-SQL-SouthEastAsia", "managedInstance", AuthenticationName.DEFAULT,
            com.azure.core.util.Context.NONE);
    }
}
