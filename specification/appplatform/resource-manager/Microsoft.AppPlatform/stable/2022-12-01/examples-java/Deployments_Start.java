import com.azure.core.util.Context;

/** Samples for Deployments Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Deployments_Start.json
     */
    /**
     * Sample code: Deployments_Start.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentsStart(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getDeployments()
            .start("myResourceGroup", "myservice", "myapp", "mydeployment", Context.NONE);
    }
}
