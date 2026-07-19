
/**
 * Samples for ServerAzureADAdministrators ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/AdministratorList.json
     */
    /**
     * Sample code: Gets a list of Azure Active Directory administrator.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsAListOfAzureActiveDirectoryAdministrator(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerAzureADAdministrators().listByServer("sqlcrudtest-4799", "sqlcrudtest-6440",
            com.azure.core.util.Context.NONE);
    }
}
