
/**
 * Samples for CloudExadataInfrastructures AddStorageCapacity.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/exaInfra_addStorageCapacity.json
     */
    /**
     * Sample code: CloudExadataInfrastructures_addStorageCapacity.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void cloudExadataInfrastructuresAddStorageCapacity(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.cloudExadataInfrastructures().addStorageCapacity("rg000", "infra1", com.azure.core.util.Context.NONE);
    }
}
