import com.azure.core.util.Context;

/** Samples for Deployments Stop. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Deployments_Stop.json
     */
    /**
     * Sample code: Deployments_Stop.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentsStop(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getDeployments()
            .stop("myResourceGroup", "myservice", "myapp", "mydeployment", Context.NONE);
    }
}
