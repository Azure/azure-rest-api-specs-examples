import com.azure.core.util.Context;

/** Samples for Deployments Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Deployments_Get.json
     */
    /**
     * Sample code: Deployments_Get.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deploymentsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getDeployments()
            .getWithResponse("myResourceGroup", "myservice", "myapp", "mydeployment", Context.NONE);
    }
}
