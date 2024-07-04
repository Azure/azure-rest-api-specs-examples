
/**
 * Samples for DbServers ListByCloudExadataInfrastructure.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/dbServers_listByParent.json
     */
    /**
     * Sample code: List DbServers by Exadata Infrastructure.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        listDbServersByExadataInfrastructure(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.dbServers().listByCloudExadataInfrastructure("rg000", "infra1", com.azure.core.util.Context.NONE);
    }
}
