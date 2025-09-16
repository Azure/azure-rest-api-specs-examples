
/**
 * Samples for ContainerAppsSourceControls Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/
     * SourceControls_Delete.json
     */
    /**
     * Sample code: Delete Container App SourceControl.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        deleteContainerAppSourceControl(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsSourceControls().delete("workerapps-rg-xj", "testcanadacentral", "current",
            "githubaccesstoken", false, false, com.azure.core.util.Context.NONE);
    }
}
