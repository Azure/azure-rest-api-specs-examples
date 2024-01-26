
/**
 * Samples for ContainerRegistries Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * ContainerRegistries_Get.json
     */
    /**
     * Sample code: ContainerRegistries_Get.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void containerRegistriesGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getContainerRegistries().getWithResponse("myResourceGroup",
            "service-name", "my-container-registry", com.azure.core.util.Context.NONE);
    }
}
