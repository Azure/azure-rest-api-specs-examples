import com.azure.core.util.Context;

/** Samples for ContainerApps GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/ContainerApps_Get.json
     */
    /**
     * Sample code: Get Container App.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getContainerApp(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.containerApps().getByResourceGroupWithResponse("rg", "testcontainerApp0", Context.NONE);
    }
}
