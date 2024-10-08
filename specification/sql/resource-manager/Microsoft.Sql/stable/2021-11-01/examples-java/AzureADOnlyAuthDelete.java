
import com.azure.resourcemanager.sql.models.AuthenticationName;

/**
 * Samples for ServerAzureADOnlyAuthentications Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/AzureADOnlyAuthDelete.json
     */
    /**
     * Sample code: Deletes Azure Active Directory only authentication object.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        deletesAzureActiveDirectoryOnlyAuthenticationObject(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerAzureADOnlyAuthentications().delete("sqlcrudtest-4799",
            "sqlcrudtest-6440", AuthenticationName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
