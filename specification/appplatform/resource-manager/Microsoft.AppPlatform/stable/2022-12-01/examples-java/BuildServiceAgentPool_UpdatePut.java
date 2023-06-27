import com.azure.core.util.Context;
import com.azure.resourcemanager.appplatform.fluent.models.BuildServiceAgentPoolResourceInner;
import com.azure.resourcemanager.appplatform.models.BuildServiceAgentPoolProperties;
import com.azure.resourcemanager.appplatform.models.BuildServiceAgentPoolSizeProperties;

/** Samples for BuildServiceAgentPool UpdatePut. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/BuildServiceAgentPool_UpdatePut.json
     */
    /**
     * Sample code: BuildServiceAgentPool_UpdatePut.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void buildServiceAgentPoolUpdatePut(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getBuildServiceAgentPools()
            .updatePut(
                "myResourceGroup",
                "myservice",
                "default",
                "default",
                new BuildServiceAgentPoolResourceInner()
                    .withProperties(
                        new BuildServiceAgentPoolProperties()
                            .withPoolSize(new BuildServiceAgentPoolSizeProperties().withName("S3"))),
                Context.NONE);
    }
}
