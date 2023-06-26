import com.azure.core.util.Context;

/** Samples for Deployments DisableRemoteDebugging. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Deployments_DisableRemoteDebugging.json
     */
    /**
     * Sample code: Deployments_DisableRemoteDebugging.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentsDisableRemoteDebugging(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getDeployments()
            .disableRemoteDebugging("myResourceGroup", "myservice", "myapp", "mydeployment", Context.NONE);
    }
}
