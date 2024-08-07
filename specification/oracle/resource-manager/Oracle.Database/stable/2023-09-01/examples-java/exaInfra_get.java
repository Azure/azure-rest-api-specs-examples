
/**
 * Samples for CloudExadataInfrastructures GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/oracle/resource-manager/Oracle.Database/stable/2023-09-01/examples/exaInfra_get.json
     */
    /**
     * Sample code: Get Exadata Infrastructure.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        getExadataInfrastructure(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.cloudExadataInfrastructures().getByResourceGroupWithResponse("rg000", "infra1",
            com.azure.core.util.Context.NONE);
    }
}
