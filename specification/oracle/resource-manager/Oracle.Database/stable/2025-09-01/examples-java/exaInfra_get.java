
/**
 * Samples for CloudExadataInfrastructures GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/exaInfra_get.json
     */
    /**
     * Sample code: CloudExadataInfrastructures_get.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        cloudExadataInfrastructuresGet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.cloudExadataInfrastructures().getByResourceGroupWithResponse("rg000", "infra1",
            com.azure.core.util.Context.NONE);
    }
}
