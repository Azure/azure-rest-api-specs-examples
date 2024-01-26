
/**
 * Samples for BuildServiceBuilder Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * BuildServiceBuilder_Get.json
     */
    /**
     * Sample code: BuildServiceBuilder_Get.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void buildServiceBuilderGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getBuildServiceBuilders().getWithResponse("myResourceGroup",
            "myservice", "default", "mybuilder", com.azure.core.util.Context.NONE);
    }
}
