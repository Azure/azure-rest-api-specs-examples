
/**
 * Samples for NetworkWatchers ConnectionAnalyzersDelete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/NetworkWatcherConnectionAnalyzerDelete.json
     */
    /**
     * Sample code: Delete connection analyzer.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void deleteConnectionAnalyzer(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkWatchers().connectionAnalyzersDelete("connectionAnalyzerRG", "nw1", "ca1",
            com.azure.core.util.Context.NONE);
    }
}
