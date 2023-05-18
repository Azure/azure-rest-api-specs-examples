import com.azure.resourcemanager.mysqlflexibleserver.models.AdministratorName;

/** Samples for AzureADAdministrators Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/AAD/preview/2021-12-01-preview/examples/AzureADAdministratorDelete.json
     */
    /**
     * Sample code: Delete an azure ad administrator.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void deleteAnAzureAdAdministrator(
        com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager
            .azureADAdministrators()
            .delete("testrg", "mysqltestsvc4", AdministratorName.ACTIVE_DIRECTORY, com.azure.core.util.Context.NONE);
    }
}
