/** Samples for ContentKeyPolicies List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-list-by-lastModified.json
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
            .list("contosorg", "contosomedia", null, null, "properties/lastModified", com.azure.core.util.Context.NONE);
    }
}
