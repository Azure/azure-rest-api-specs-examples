
/**
 * Samples for BuildsByBuilderResource List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/Builds_ListByBuilderResource
     * .json
     */
    /**
     * Sample code: Builds_ListByBuilderResource_0.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        buildsListByBuilderResource0(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.buildsByBuilderResources().list("rg", "testBuilder", com.azure.core.util.Context.NONE);
    }
}
