
/**
 * Samples for BuildServiceBuilder Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * BuildServiceBuilder_Delete.json
     */
    /**
     * Sample code: BuildServiceBuilder_Delete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void buildServiceBuilderDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getBuildServiceBuilders().delete("myResourceGroup",
            "myservice", "default", "mybuilder", com.azure.core.util.Context.NONE);
    }
}
