import com.azure.core.util.Context;

/** Samples for AccountFilters Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/accountFilters-get-by-name.json
     */
    /**
     * Sample code: Get an Account Filter by name.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getAnAccountFilterByName(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager.accountFilters().getWithResponse("contoso", "contosomedia", "accountFilterWithTrack", Context.NONE);
    }
}
