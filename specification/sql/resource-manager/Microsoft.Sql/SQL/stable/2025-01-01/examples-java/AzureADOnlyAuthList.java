
/**
 * Samples for ServerAzureADOnlyAuthentications ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/AzureADOnlyAuthList.json
     */
    /**
     * Sample code: Gets a list of Azure Active Directory only authentication object.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsAListOfAzureActiveDirectoryOnlyAuthenticationObject(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerAzureADOnlyAuthentications().listByServer("sqlcrudtest-4799",
            "sqlcrudtest-6440", com.azure.core.util.Context.NONE);
    }
}
