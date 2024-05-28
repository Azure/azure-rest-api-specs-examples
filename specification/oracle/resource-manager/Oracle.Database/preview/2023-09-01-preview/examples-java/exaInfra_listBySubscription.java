
/**
 * Samples for CloudExadataInfrastructures List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/
     * exaInfra_listBySubscription.json
     */
    /**
     * Sample code: List Exadata Infrastructure by subscription.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void listExadataInfrastructureBySubscription(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.cloudExadataInfrastructures().list(com.azure.core.util.Context.NONE);
    }
}
