
/**
 * Samples for ContainerApps ListCustomHostnameAnalysis.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/ContainerApps_ListCustomHostNameAnalysis.json
     */
    /**
     * Sample code: Analyze Custom Hostname.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void analyzeCustomHostname(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerApps().listCustomHostnameAnalysisWithResponse("rg", "testcontainerApp0", "my.name.corp",
            com.azure.core.util.Context.NONE);
    }
}
