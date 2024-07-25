
/**
 * Samples for DotNetComponents Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2024-02-02-preview/examples/DotNetComponents_Get.json
     */
    /**
     * Sample code: Get .NET Component.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void getNETComponent(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.dotNetComponents().getWithResponse("examplerg", "myenvironment", "mydotnetcomponent",
            com.azure.core.util.Context.NONE);
    }
}
