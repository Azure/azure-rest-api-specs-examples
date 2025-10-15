
/**
 * Samples for ExascaleDbStorageVaults Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/ExascaleDbStorageVaults_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: ExascaleDbStorageVaults_Delete_MaximumSet.
     * 
     * @param manager Entry point to OracleDatabaseManager.
     */
    public static void exascaleDbStorageVaultsDeleteMaximumSet(
        com.azure.resourcemanager.oracledatabase.OracleDatabaseManager manager) {
        manager.exascaleDbStorageVaults().delete("rgopenapi", "storagevault1", com.azure.core.util.Context.NONE);
    }
}
