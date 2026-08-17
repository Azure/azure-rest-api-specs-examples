
/**
 * Samples for NetworkWatchers ConnectionAnalyzersQuery.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/NetworkWatcherConnectionAnalyzerQuery.json
     */
    /**
     * Sample code: Query connection analyzer.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void queryConnectionAnalyzer(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkWatchers().connectionAnalyzersQuery("connectionAnalyzerRG", "nw1", "ca1",
            com.azure.core.util.Context.NONE);
    }
}
