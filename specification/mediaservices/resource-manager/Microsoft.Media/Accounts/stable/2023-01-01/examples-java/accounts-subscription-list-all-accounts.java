/** Samples for Mediaservices List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Accounts/stable/2023-01-01/examples/accounts-subscription-list-all-accounts.json
     */
    /**
     * Sample code: List all Media Services accounts.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listAllMediaServicesAccounts(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.mediaservices().list(com.azure.core.util.Context.NONE);
    }
}
