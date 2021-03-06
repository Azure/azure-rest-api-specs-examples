import com.azure.core.util.Context;

/** Samples for Configurations ListByServer. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2021-05-01/examples/ConfigurationsListByServer.json
     */
    /**
     * Sample code: List all configurations for a server.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void listAllConfigurationsForAServer(
        com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.configurations().listByServer("testrg", "mysqltestserver", Context.NONE);
    }
}
