
/**
 * Samples for PrivateLinkServices ListAutoApprovedPrivateLinkServices.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/AutoApprovedPrivateLinkServicesGet.json
     */
    /**
     * Sample code: Get list of private link service id that can be linked to a private end point with auto approved.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void getListOfPrivateLinkServiceIdThatCanBeLinkedToAPrivateEndPointWithAutoApproved(
        com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getPrivateLinkServices().listAutoApprovedPrivateLinkServices("regionName",
            com.azure.core.util.Context.NONE);
    }
}
