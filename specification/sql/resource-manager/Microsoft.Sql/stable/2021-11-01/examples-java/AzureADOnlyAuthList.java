
/**
 * Samples for ServerAzureADOnlyAuthentications ListByServer.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/AzureADOnlyAuthList.json
     */
    /**
     * Sample code: Gets a list of Azure Active Directory only authentication object.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getsAListOfAzureActiveDirectoryOnlyAuthenticationObject(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getServerAzureADOnlyAuthentications()
            .listByServer("sqlcrudtest-4799", "sqlcrudtest-6440", com.azure.core.util.Context.NONE);
    }
}
