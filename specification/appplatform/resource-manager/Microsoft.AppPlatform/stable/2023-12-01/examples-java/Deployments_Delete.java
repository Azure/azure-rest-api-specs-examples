
/**
 * Samples for Deployments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/Deployments_Delete.
     * json
     */
    /**
     * Sample code: Deployments_Delete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getDeployments().delete("myResourceGroup", "myservice",
            "myapp", "mydeployment", com.azure.core.util.Context.NONE);
    }
}
