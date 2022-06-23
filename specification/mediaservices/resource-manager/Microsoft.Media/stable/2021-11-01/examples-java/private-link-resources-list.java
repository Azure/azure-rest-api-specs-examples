import com.azure.core.util.Context;

/** Samples for PrivateLinkResources List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/private-link-resources-list.json
     */
    /**
     * Sample code: Get list of all group IDs.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getListOfAllGroupIDs(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.privateLinkResources().listWithResponse("contoso", "contososports", Context.NONE);
    }
}
