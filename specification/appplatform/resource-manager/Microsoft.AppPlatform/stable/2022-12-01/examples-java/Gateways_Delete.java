import com.azure.core.util.Context;

/** Samples for Gateways Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-12-01/examples/Gateways_Delete.json
     */
    /**
     * Sample code: Gateways_Delete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void gatewaysDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .springServices()
            .manager()
            .serviceClient()
            .getGateways()
            .delete("myResourceGroup", "myservice", "default", Context.NONE);
    }
}
