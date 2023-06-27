import com.azure.core.util.Context;

/** Samples for Deployments ListForCluster. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Deployments_ListForCluster.json
     */
    /**
     * Sample code: Deployments_ListForCluster.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentsListForCluster(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getDeployments()
            .listForCluster("myResourceGroup", "myservice", null, Context.NONE);
    }
}
