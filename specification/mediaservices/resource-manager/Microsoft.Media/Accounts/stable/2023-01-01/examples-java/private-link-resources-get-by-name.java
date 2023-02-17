/** Samples for PrivateLinkResources Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2023-01-01/examples/private-link-resources-get-by-name.json
     */
    /**
     * Sample code: Get details of a group ID.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getDetailsOfAGroupID(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .privateLinkResources()
            .getWithResponse("contosorg", "contososports", "keydelivery", com.azure.core.util.Context.NONE);
    }
}
