
/**
 * Samples for ExascaleDbStorageVaults GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ExascaleDbStorageVaults_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExascaleDbStorageVaults_Get_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void
        exascaleDbStorageVaultsGetMaximumSet(com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.exascaleDbStorageVaults().getByResourceGroupWithResponse("rgopenapi", "storagevault1",
            com.azure.core.util.Context.NONE);
    }
}
