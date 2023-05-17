/** Samples for ContainerApps ListCustomHostnameAnalysis. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/ContainerApps_ListCustomHostNameAnalysis.json
     */
    /**
     * Sample code: Analyze Custom Hostname.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void analyzeCustomHostname(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .containerApps()
            .listCustomHostnameAnalysisWithResponse(
                "rg", "testcontainerApp0", "my.name.corp", com.azure.core.util.Context.NONE);
    }
}
