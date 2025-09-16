
/**
 * Samples for ContainerApps ListCustomHostnameAnalysis.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/
     * ContainerApps_ListCustomHostNameAnalysis.json
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
