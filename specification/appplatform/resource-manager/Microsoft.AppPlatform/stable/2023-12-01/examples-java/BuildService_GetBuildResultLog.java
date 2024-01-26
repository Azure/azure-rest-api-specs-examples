
/**
 * Samples for BuildService GetBuildResultLog.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * BuildService_GetBuildResultLog.json
     */
    /**
     * Sample code: BuildService_GetBuildResultLog.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void buildServiceGetBuildResultLog(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getBuildServices().getBuildResultLogWithResponse(
            "myResourceGroup", "myservice", "default", "mybuild", "123", com.azure.core.util.Context.NONE);
    }
}
