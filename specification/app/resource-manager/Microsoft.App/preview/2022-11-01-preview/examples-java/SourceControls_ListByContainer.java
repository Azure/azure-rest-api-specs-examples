/** Samples for ContainerAppsSourceControls ListByContainerApp. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/SourceControls_ListByContainer.json
     */
    /**
     * Sample code: List App's Source Controls.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listAppSSourceControls(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .containerAppsSourceControls()
            .listByContainerApp("workerapps-rg-xj", "testcanadacentral", com.azure.core.util.Context.NONE);
    }
}
