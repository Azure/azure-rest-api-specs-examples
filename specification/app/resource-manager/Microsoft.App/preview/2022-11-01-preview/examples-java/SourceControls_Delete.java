/** Samples for ContainerAppsSourceControls Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/SourceControls_Delete.json
     */
    /**
     * Sample code: Delete Container App SourceControl.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deleteContainerAppSourceControl(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .containerAppsSourceControls()
            .delete("workerapps-rg-xj", "testcanadacentral", "current", com.azure.core.util.Context.NONE);
    }
}
