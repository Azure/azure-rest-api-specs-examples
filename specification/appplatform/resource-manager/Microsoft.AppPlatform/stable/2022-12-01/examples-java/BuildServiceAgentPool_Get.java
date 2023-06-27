import com.azure.core.util.Context;

/** Samples for BuildServiceAgentPool Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/BuildServiceAgentPool_Get.json
     */
    /**
     * Sample code: BuildServiceAgentPool_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void buildServiceAgentPoolGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getBuildServiceAgentPools()
            .getWithResponse("myResourceGroup", "myservice", "default", "default", Context.NONE);
    }
}
