
/**
 * Samples for NetworkWatchers ConnectionAnalyzersList.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/NetworkWatcherConnectionAnalyzerList.json
     */
    /**
     * Sample code: List connection analyzers.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listConnectionAnalyzers(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getNetworkWatchers().connectionAnalyzersList("connectionAnalyzerRG", "nw1",
            com.azure.core.util.Context.NONE);
    }
}
