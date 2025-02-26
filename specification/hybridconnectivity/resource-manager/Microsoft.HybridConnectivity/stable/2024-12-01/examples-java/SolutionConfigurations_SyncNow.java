
/**
 * Samples for SolutionConfigurations SyncNow.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/SolutionConfigurations_SyncNow.json
     */
    /**
     * Sample code: SolutionConfigurations_SyncNow.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void
        solutionConfigurationsSyncNow(com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager.solutionConfigurations().syncNow("ymuj", "tks", com.azure.core.util.Context.NONE);
    }
}
