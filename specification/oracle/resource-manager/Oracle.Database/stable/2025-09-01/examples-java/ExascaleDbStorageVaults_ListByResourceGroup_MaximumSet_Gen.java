
/**
 * Samples for ExascaleDbStorageVaults ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ExascaleDbStorageVaults_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExascaleDbStorageVaults_ListByResourceGroup_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void exascaleDbStorageVaultsListByResourceGroupMaximumSet(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.exascaleDbStorageVaults().listByResourceGroup("rgopenapi", com.azure.core.util.Context.NONE);
    }
}
