import com.azure.core.util.Context;

/** Samples for Mediaservices ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-06-01/examples/accounts-list-all-accounts.json
     */
    /**
     * Sample code: List all Media Services accounts.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listAllMediaServicesAccounts(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.mediaservices().listByResourceGroup("contoso", Context.NONE);
    }
}
