import com.azure.core.util.Context;

/** Samples for Deployments List. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Deployments_List.json
     */
    /**
     * Sample code: Deployments_List.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getDeployments()
            .list("myResourceGroup", "myservice", "myapp", null, Context.NONE);
    }
}
