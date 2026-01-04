
/**
 * Samples for Registries GetBuildSourceUploadUrl.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/RegistryTasks/preview/2019-06-01-
     * preview/examples/RegistriesGetBuildSourceUploadUrl.json
     */
    /**
     * Sample code: Registries_GetBuildSourceUploadUrl.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registriesGetBuildSourceUploadUrl(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.containerRegistries().manager().serviceClient().getRegistries()
            .getBuildSourceUploadUrlWithResponse("myResourceGroup", "myRegistry", com.azure.core.util.Context.NONE);
    }
}
