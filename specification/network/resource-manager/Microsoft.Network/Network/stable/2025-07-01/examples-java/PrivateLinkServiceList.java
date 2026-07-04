
/**
 * Samples for PrivateLinkServices ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/PrivateLinkServiceList.json
     */
    /**
     * Sample code: List private link service in resource group.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void listPrivateLinkServiceInResourceGroup(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPrivateLinkServices().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
