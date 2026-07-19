
import com.azure.resourcemanager.sql.fluent.models.ServerAzureADOnlyAuthenticationInner;
import com.azure.resourcemanager.sql.models.AuthenticationName;

/**
 * Samples for ServerAzureADOnlyAuthentications CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/AzureADOnlyAuthCreateOrUpdate.json
     */
    /**
     * Sample code: Creates or updates Azure Active Directory only authentication object.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createsOrUpdatesAzureActiveDirectoryOnlyAuthenticationObject(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerAzureADOnlyAuthentications().createOrUpdate("sqlcrudtest-4799",
            "sqlcrudtest-6440", AuthenticationName.DEFAULT,
            new ServerAzureADOnlyAuthenticationInner().withAzureADOnlyAuthentication(false),
            com.azure.core.util.Context.NONE);
    }
}
