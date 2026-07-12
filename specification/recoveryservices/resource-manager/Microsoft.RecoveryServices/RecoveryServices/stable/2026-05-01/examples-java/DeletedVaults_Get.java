
/**
 * Samples for DeletedVaults Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/DeletedVaults_Get.json
     */
    /**
     * Sample code: Gets Deleted vault.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void getsDeletedVault(com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.deletedVaults().getWithResponse("westus", "swaggerExample", com.azure.core.util.Context.NONE);
    }
}
