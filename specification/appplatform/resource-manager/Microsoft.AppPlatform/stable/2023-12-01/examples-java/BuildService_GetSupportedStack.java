
/**
 * Samples for BuildService GetSupportedStack.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * BuildService_GetSupportedStack.json
     */
    /**
     * Sample code: BuildService_GetSupportedStack.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void buildServiceGetSupportedStack(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getBuildServices().getSupportedStackWithResponse(
            "myResourceGroup", "myservice", "default", "io.buildpacks.stacks.bionic-base",
            com.azure.core.util.Context.NONE);
    }
}
