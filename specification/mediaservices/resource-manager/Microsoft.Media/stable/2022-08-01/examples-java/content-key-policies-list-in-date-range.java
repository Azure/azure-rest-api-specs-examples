import com.azure.core.util.Context;

/** Samples for ContentKeyPolicies List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2022-08-01/examples/content-key-policies-list-in-date-range.json
     */
    /**
     * Sample code: Lists Content Key Policies with created and last modified filters.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listsContentKeyPoliciesWithCreatedAndLastModifiedFilters(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .contentKeyPolicies()
            .list(
                "contoso",
                "contosomedia",
                "properties/lastModified gt 2016-06-01 and properties/created lt 2013-07-01",
                null,
                null,
                Context.NONE);
    }
}
