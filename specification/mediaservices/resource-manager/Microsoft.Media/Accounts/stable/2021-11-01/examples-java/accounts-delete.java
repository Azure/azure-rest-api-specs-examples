/** Samples for Mediaservices Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2021-11-01/examples/accounts-delete.json
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
            .deleteByResourceGroupWithResponse("contoso", "contososports", com.azure.core.util.Context.NONE);
    }
}
