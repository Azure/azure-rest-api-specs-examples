import com.azure.core.util.Context;

/** Samples for ContainerAppsSourceControls Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-01-01-preview/examples/SourceControls_Get.json
     */
    /**
     * Sample code: Get Container App's SourceControl.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getContainerAppSSourceControl(
        com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .containerAppsSourceControls()
            .getWithResponse("workerapps-rg-xj", "testcanadacentral", "current", Context.NONE);
    }
}
