import com.azure.core.util.Context;

/** Samples for Transforms List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/transforms-list-all-filter-by-lastmodified.json
     */
    /**
     * Sample code: Lists the Transforms filter by lastmodified.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listsTheTransformsFilterByLastmodified(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .transforms()
            .list(
                "contosoresources",
                "contosomedia",
                "properties/lastmodified gt 2021-11-01T00:00:00.0000000Z and properties/lastmodified le"
                    + " 2021-11-01T00:00:10.0000000Z",
                "properties/lastmodified desc",
                Context.NONE);
    }
}
