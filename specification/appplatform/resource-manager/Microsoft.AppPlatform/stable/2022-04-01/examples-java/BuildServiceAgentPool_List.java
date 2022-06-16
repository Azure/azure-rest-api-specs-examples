import com.azure.core.util.Context;

/** Samples for BuildServiceAgentPool List. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/BuildServiceAgentPool_List.json
     */
    /**
     * Sample code: BuildServiceAgentPool_List.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void buildServiceAgentPoolList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getBuildServiceAgentPools()
            .list("myResourceGroup", "myservice", "default", Context.NONE);
    }
}
