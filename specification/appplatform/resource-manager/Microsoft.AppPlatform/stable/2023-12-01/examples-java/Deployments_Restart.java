
/**
 * Samples for Deployments Restart.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/Deployments_Restart.
     * json
     */
    /**
     * Sample code: Deployments_Restart.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentsRestart(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getDeployments().restart("myResourceGroup", "myservice",
            "myapp", "mydeployment", com.azure.core.util.Context.NONE);
    }
}
