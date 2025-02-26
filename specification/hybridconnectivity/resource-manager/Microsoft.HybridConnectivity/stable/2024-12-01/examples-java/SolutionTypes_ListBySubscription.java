
/**
 * Samples for SolutionTypes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-12-01/SolutionTypes_ListBySubscription.json
     */
    /**
     * Sample code: SolutionTypes_ListBySubscription.
     * 
     * @param manager Entry point to HybridConnectivityManager.
     */
    public static void solutionTypesListBySubscription(
        com.azure.resourcemanager.hybridconnectivity.HybridConnectivityManager manager) {
        manager.solutionTypes().list(com.azure.core.util.Context.NONE);
    }
}
