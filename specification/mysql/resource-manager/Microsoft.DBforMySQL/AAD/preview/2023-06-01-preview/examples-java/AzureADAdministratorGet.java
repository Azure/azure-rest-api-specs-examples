
import com.azure.resourcemanager.mysqlflexibleserver.models.AdministratorName;

/**
 * Samples for AzureADAdministrators Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/AAD/preview/2023-06-01-preview/examples/
     * AzureADAdministratorGet.json
     */
    /**
     * Sample code: Get an azure ad administrator.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void getAnAzureAdAdministrator(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.azureADAdministrators().getWithResponse("testrg", "mysqltestsvc4", AdministratorName.ACTIVE_DIRECTORY,
            com.azure.core.util.Context.NONE);
    }
}
