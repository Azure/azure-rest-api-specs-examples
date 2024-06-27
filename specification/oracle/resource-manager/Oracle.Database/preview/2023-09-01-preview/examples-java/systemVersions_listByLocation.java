
/**
 * Samples for SystemVersions ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/
     * systemVersions_listByLocation.json
     */
    /**
     * Sample code: List Exadata System Versions by the provided filter.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void listExadataSystemVersionsByTheProvidedFilter(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.systemVersions().listByLocation("eastus", com.azure.core.util.Context.NONE);
    }

    /*
     * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/
     * systemVersions_listByLocation.json
     */
    /**
     * Sample code: systemVersions_listByLocation.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        systemVersionsListByLocation(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.systemVersions().listByLocation("eastus", com.azure.core.util.Context.NONE);
    }
}
