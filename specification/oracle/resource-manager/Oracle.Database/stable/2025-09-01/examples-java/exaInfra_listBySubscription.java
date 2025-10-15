
/**
 * Samples for CloudExadataInfrastructures List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/exaInfra_listBySubscription.json
     */
    /**
     * Sample code: CloudExadataInfrastructures_listBySubscription.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void cloudExadataInfrastructuresListBySubscription(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.cloudExadataInfrastructures().list(com.azure.core.util.Context.NONE);
    }
}
