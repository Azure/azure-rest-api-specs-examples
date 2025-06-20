
/**
 * Samples for DbServers ListByCloudExadataInfrastructure.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01/dbServers_listByParent.json
     */
    /**
     * Sample code: DbServers_listByCloudExadataInfrastructure.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void dbServersListByCloudExadataInfrastructure(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dbServers().listByCloudExadataInfrastructure("rg000", "infra1", com.azure.core.util.Context.NONE);
    }
}
