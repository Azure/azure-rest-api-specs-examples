import com.azure.core.util.Context;

/** Samples for Transforms List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/transforms-list-all-filter-by-name.json
     */
    /**
     * Sample code: Lists the Transforms filter by name.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listsTheTransformsFilterByName(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .transforms()
            .list(
                "contosoresources",
                "contosomedia",
                "(name eq 'sampleEncode') or (name eq 'sampleEncodeAndVideoIndex')",
                "name desc",
                Context.NONE);
    }
}
