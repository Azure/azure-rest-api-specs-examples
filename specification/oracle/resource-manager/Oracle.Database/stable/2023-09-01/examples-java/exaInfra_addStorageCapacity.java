
/**
 * Samples for CloudExadataInfrastructures AddStorageCapacity.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/exaInfra_addStorageCapacity.json
     */
    /**
     * Sample code: Perform add storage capacity on exadata infra.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void performAddStorageCapacityOnExadataInfra(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.cloudExadataInfrastructures().addStorageCapacity("rg000", "infra1", com.azure.core.util.Context.NONE);
    }
}
