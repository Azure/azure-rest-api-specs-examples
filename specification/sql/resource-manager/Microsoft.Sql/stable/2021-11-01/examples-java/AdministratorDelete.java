
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.models.AdministratorName;

/** Samples for ServerAzureADAdministrators Delete. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/AdministratorDelete.json
     */
    /**
     * Sample code: Delete Azure Active Directory administrator.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAzureActiveDirectoryAdministrator(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerAzureADAdministrators().delete("sqlcrudtest-4799",
            "sqlcrudtest-6440", AdministratorName.ACTIVE_DIRECTORY, Context.NONE);
    }
}
