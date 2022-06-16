import com.azure.core.util.Context;

/** Samples for ContainerAppsSourceControls Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/SourceControls_Delete.json
     */
    /**
     * Sample code: Delete Container App SourceControl.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deleteContainerAppSourceControl(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerAppsSourceControls().delete("workerapps-rg-xj", "testcanadacentral", "current", Context.NONE);
    }
}
