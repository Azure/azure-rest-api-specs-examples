import com.azure.core.util.Context;

/** Samples for PrivateLinkServices GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/PrivateLinkServiceGet.json
     */
    /**
     * Sample code: Get private link service.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getPrivateLinkService(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getPrivateLinkServices()
            .getByResourceGroupWithResponse("rg1", "testPls", null, Context.NONE);
    }
}
