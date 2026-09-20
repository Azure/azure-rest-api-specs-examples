
/**
 * Samples for DotNetComponents Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/DotNetComponents_Get_ServiceBind.json
     */
    /**
     * Sample code: Get .NET Component with ServiceBinds.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void
        getNETComponentWithServiceBinds(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.dotNetComponents().getWithResponse("examplerg", "myenvironment", "mydotnetcomponent",
            com.azure.core.util.Context.NONE);
    }
}
