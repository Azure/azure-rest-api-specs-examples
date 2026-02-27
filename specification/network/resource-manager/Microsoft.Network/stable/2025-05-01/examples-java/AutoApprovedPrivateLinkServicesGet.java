
/**
 * Samples for PrivateLinkServices ListAutoApprovedPrivateLinkServices.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2025-05-01/examples/
     * AutoApprovedPrivateLinkServicesGet.json
     */
    /**
     * Sample code: Get list of private link service id that can be linked to a private end point with auto approved.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getListOfPrivateLinkServiceIdThatCanBeLinkedToAPrivateEndPointWithAutoApproved(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.networks().manager().serviceClient().getPrivateLinkServices()
            .listAutoApprovedPrivateLinkServices("regionName", com.azure.core.util.Context.NONE);
    }
}
