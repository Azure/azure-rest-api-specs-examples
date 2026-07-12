
/**
 * Samples for Usages ListByVaults.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/ListUsages.json
     */
    /**
     * Sample code: Gets vault usages.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void getsVaultUsages(com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.usages().listByVaults("Default-RecoveryServices-ResourceGroup", "swaggerExample",
            com.azure.core.util.Context.NONE);
    }
}
