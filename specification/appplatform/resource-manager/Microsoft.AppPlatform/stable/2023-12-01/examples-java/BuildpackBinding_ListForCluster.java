
/**
 * Samples for BuildpackBinding ListForCluster.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/
     * BuildpackBinding_ListForCluster.json
     */
    /**
     * Sample code: BuildpackBinding_ListForCluster.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void buildpackBindingListForCluster(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getBuildpackBindings().listForCluster("myResourceGroup",
            "myservice", com.azure.core.util.Context.NONE);
    }
}
