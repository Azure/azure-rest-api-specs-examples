
/**
 * Samples for DotNetComponents Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/preview/2023-11-02-preview/examples/DotNetComponents_Delete.json
     */
    /**
     * Sample code: Delete .NET Component.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void deleteNETComponent(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.dotNetComponents().delete("examplerg", "myenvironment", "mydotnetcomponent",
            com.azure.core.util.Context.NONE);
    }
}
