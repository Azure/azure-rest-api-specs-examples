
/**
 * Samples for BuildService GetBuildResult.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * BuildService_GetBuildResult.json
     */
    /**
     * Sample code: BuildService_GetBuildResult.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void buildServiceGetBuildResult(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getBuildServices().getBuildResultWithResponse(
            "myResourceGroup", "myservice", "default", "mybuild", "123", com.azure.core.util.Context.NONE);
    }
}
