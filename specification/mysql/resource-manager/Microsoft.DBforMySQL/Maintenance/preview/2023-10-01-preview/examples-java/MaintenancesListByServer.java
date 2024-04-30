
/**
 * Samples for Maintenances List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/mysql/resource-manager/Microsoft.DBforMySQL/Maintenance/preview/2023-10-01-preview/examples/
     * MaintenancesListByServer.json
     */
    /**
     * Sample code: List maintenances on a server.
     * 
     * @param manager Entry point to MySqlManager.
     */
    public static void listMaintenancesOnAServer(com.azure.resourcemanager.mysqlflexibleserver.MySqlManager manager) {
        manager.maintenances().list("TestGroup", "testserver", com.azure.core.util.Context.NONE);
    }
}
