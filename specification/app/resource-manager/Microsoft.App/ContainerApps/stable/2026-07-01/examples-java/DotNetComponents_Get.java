
/**
 * Samples for DotNetComponents Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/DotNetComponents_Get.json
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
