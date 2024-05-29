
/**
 * Samples for CloudExadataInfrastructures ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/oracle/resource-manager/Oracle.Database/preview/2023-09-01-preview/examples/
     * exaInfra_listByResourceGroup.json
     */
    /**
     * Sample code: List Exadata Infrastructure by resource group.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void listExadataInfrastructureByResourceGroup(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.cloudExadataInfrastructures().listByResourceGroup("rg000", com.azure.core.util.Context.NONE);
    }
}
