
/**
 * Samples for CloudExadataInfrastructures ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/exaInfra_listByResourceGroup.json
     */
    /**
     * Sample code: CloudExadataInfrastructures_listByResourceGroup.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void cloudExadataInfrastructuresListByResourceGroup(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.cloudExadataInfrastructures().listByResourceGroup("rg000", com.azure.core.util.Context.NONE);
    }
}
