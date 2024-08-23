
import com.azure.resourcemanager.sql.models.AuthenticationName;

/**
 * Samples for ManagedInstanceAzureADOnlyAuthentications Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedInstanceAzureADOnlyAuthDelete.
     * json
     */
    /**
     * Sample code: Deletes Azure Active Directory only authentication object.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        deletesAzureActiveDirectoryOnlyAuthenticationObject(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstanceAzureADOnlyAuthentications().delete(
            "Default-SQL-SouthEastAsia", "managedInstance", AuthenticationName.DEFAULT,
            com.azure.core.util.Context.NONE);
    }
}
