/** Samples for ContainerAppsDiagnostics ListDetectors. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/ContainerAppsDiagnostics_List.json
     */
    /**
     * Sample code: Get the list of available diagnostics for a given Container App.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getTheListOfAvailableDiagnosticsForAGivenContainerApp(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .containerAppsDiagnostics()
            .listDetectors("mikono-workerapp-test-rg", "mikono-capp-stage1", com.azure.core.util.Context.NONE);
    }
}
