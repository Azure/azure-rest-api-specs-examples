import com.azure.core.util.Context;

/** Samples for Deployments GetLogFileUrl. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Deployments_GetLogFileUrl.json
     */
    /**
     * Sample code: Deployments_GetLogFileUrl.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentsGetLogFileUrl(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getDeployments()
            .getLogFileUrlWithResponse("myResourceGroup", "myservice", "myapp", "mydeployment", Context.NONE);
    }
}
