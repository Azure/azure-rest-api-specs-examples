
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ServerAzureADOnlyAuthenticationInner;
import com.azure.resourcemanager.sql.models.AuthenticationName;

/** Samples for ServerAzureADOnlyAuthentications CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/AzureADOnlyAuthCreateOrUpdate.json
     */
    /**
     * Sample code: Creates or updates Azure Active Directory only authentication object.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createsOrUpdatesAzureActiveDirectoryOnlyAuthenticationObject(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerAzureADOnlyAuthentications().createOrUpdate(
            "sqlcrudtest-4799", "sqlcrudtest-6440", AuthenticationName.DEFAULT,
            new ServerAzureADOnlyAuthenticationInner().withAzureADOnlyAuthentication(false), Context.NONE);
    }
}
