import com.azure.core.util.Context;
import com.azure.resourcemanager.network.models.CheckPrivateLinkServiceVisibilityRequest;

/** Samples for PrivateLinkServices CheckPrivateLinkServiceVisibilityByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/CheckPrivateLinkServiceVisibilityByResourceGroup.json
     */
    /**
     * Sample code: Check private link service visibility.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void checkPrivateLinkServiceVisibility(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .networks()
            .manager()
            .serviceClient()
            .getPrivateLinkServices()
            .checkPrivateLinkServiceVisibilityByResourceGroup(
                "westus",
                "rg1",
                new CheckPrivateLinkServiceVisibilityRequest()
                    .withPrivateLinkServiceAlias("mypls.00000000-0000-0000-0000-000000000000.azure.privatelinkservice"),
                Context.NONE);
    }
}
