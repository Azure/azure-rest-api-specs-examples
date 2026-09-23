
import com.azure.resourcemanager.sql.models.AuthenticationName;

/**
 * Samples for ServerAzureADOnlyAuthentications Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/AzureADOnlyAuthDelete.json
     */
    /**
     * Sample code: Deletes Azure Active Directory only authentication object.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        deletesAzureActiveDirectoryOnlyAuthenticationObject(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerAzureADOnlyAuthentications().delete("sqlcrudtest-4799", "sqlcrudtest-6440",
            AuthenticationName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
