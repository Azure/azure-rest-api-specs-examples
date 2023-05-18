/** Samples for AzureADAdministrators ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/AAD/preview/2021-12-01-preview/examples/AzureADAdministratorsListByServer.json
     */
    /**
     * Sample code: List Azure AD administrators in a server.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void listAzureADAdministratorsInAServer(
        com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.azureADAdministrators().listByServer("testrg", "mysqltestsvc4", com.azure.core.util.Context.NONE);
    }
}
