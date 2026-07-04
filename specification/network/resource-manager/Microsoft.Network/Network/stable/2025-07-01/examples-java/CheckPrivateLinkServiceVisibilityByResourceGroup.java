
import com.azure.resourcemanager.network.models.CheckPrivateLinkServiceVisibilityRequest;

/**
 * Samples for PrivateLinkServices CheckPrivateLinkServiceVisibilityByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/CheckPrivateLinkServiceVisibilityByResourceGroup.json
     */
    /**
     * Sample code: Check private link service visibility.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void checkPrivateLinkServiceVisibility(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPrivateLinkServices().checkPrivateLinkServiceVisibilityByResourceGroup("westus",
            "rg1",
            new CheckPrivateLinkServiceVisibilityRequest()
                .withPrivateLinkServiceAlias("mypls.00000000-0000-0000-0000-000000000000.azure.privatelinkservice"),
            com.azure.core.util.Context.NONE);
    }
}
