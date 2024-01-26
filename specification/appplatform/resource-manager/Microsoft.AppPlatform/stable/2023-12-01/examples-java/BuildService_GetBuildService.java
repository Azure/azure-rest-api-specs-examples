
/**
 * Samples for BuildService GetBuildService.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * BuildService_GetBuildService.json
     */
    /**
     * Sample code: BuildService_GetBuildService.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void buildServiceGetBuildService(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getBuildServices()
            .getBuildServiceWithResponse("myResourceGroup", "myservice", "default", com.azure.core.util.Context.NONE);
    }
}
