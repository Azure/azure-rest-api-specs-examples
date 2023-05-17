/** Samples for ManagedEnvironments GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-11-01-preview/examples/ManagedEnvironments_Get.json
     */
    /**
     * Sample code: Get environments by name.
     *
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getEnvironmentsByName(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager
            .managedEnvironments()
            .getByResourceGroupWithResponse("examplerg", "jlaw-demo1", com.azure.core.util.Context.NONE);
    }
}
