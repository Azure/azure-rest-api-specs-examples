/** Samples for Mediaservices Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2023-01-01/examples/accounts-delete.json
     */
    /**
     * Sample code: Delete a Media Services account.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void deleteAMediaServicesAccount(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .mediaservices()
            .deleteByResourceGroupWithResponse("contosorg", "contososports", com.azure.core.util.Context.NONE);
    }
}
