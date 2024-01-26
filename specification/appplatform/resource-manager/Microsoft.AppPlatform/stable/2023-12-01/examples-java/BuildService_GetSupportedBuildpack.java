
/**
 * Samples for BuildService GetSupportedBuildpack.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * BuildService_GetSupportedBuildpack.json
     */
    /**
     * Sample code: BuildService_GetSupportedBuildpack.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void buildServiceGetSupportedBuildpack(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getBuildServices().getSupportedBuildpackWithResponse(
            "myResourceGroup", "myservice", "default", "tanzu-buildpacks-java-azure", com.azure.core.util.Context.NONE);
    }
}
