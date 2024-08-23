
import com.azure.resourcemanager.sql.fluent.models.ManagedInstanceAzureADOnlyAuthenticationInner;
import com.azure.resourcemanager.sql.models.AuthenticationName;

/**
 * Samples for ManagedInstanceAzureADOnlyAuthentications CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ManagedInstanceAzureADOnlyAuthCreateOrUpdate.json
     */
    /**
     * Sample code: Creates or updates Azure Active Directory only authentication object.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createsOrUpdatesAzureActiveDirectoryOnlyAuthenticationObject(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstanceAzureADOnlyAuthentications().createOrUpdate(
            "Default-SQL-SouthEastAsia", "managedInstance", AuthenticationName.DEFAULT,
            new ManagedInstanceAzureADOnlyAuthenticationInner().withAzureADOnlyAuthentication(false),
            com.azure.core.util.Context.NONE);
    }
}
