/** Samples for AccountFilters Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/accountFilters-delete.json
     */
    /**
     * Sample code: Delete an Account Filter.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void deleteAnAccountFilter(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .accountFilters()
            .deleteWithResponse(
                "contosorg", "contosomedia", "accountFilterWithTimeWindowAndTrack", com.azure.core.util.Context.NONE);
    }
}
