
/**
 * Samples for CloudExadataInfrastructures Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/exaInfra_delete.json
     */
    /**
     * Sample code: CloudExadataInfrastructures_delete.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        cloudExadataInfrastructuresDelete(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.cloudExadataInfrastructures().delete("rg000", "infra1", com.azure.core.util.Context.NONE);
    }
}
