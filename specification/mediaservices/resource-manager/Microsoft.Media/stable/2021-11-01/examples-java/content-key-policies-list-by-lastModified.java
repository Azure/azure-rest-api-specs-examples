import com.azure.core.util.Context;

/** Samples for ContentKeyPolicies List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/content-key-policies-list-by-lastModified.json
     */
    /**
     * Sample code: Lists Content Key Policies ordered by last modified.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listsContentKeyPoliciesOrderedByLastModified(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .contentKeyPolicies()
            .list("contoso", "contosomedia", null, null, "properties/lastModified", Context.NONE);
    }
}
