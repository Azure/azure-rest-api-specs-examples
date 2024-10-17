
/**
 * Samples for ContainerAppsSourceControls Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-08-02-preview/examples/SourceControls_Delete.json
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
