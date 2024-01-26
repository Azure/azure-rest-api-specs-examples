
import com.azure.resourcemanager.appplatform.models.RemoteDebuggingPayload;

/**
 * Samples for Deployments EnableRemoteDebugging.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * Deployments_EnableRemoteDebugging.json
     */
    /**
     * Sample code: Deployments_EnableRemoteDebugging.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentsEnableRemoteDebugging(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getDeployments().enableRemoteDebugging("myResourceGroup",
            "myservice", "myapp", "mydeployment", new RemoteDebuggingPayload(), com.azure.core.util.Context.NONE);
    }
}
