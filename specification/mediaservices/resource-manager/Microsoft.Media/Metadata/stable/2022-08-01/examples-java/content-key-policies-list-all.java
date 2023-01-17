/** Samples for ContentKeyPolicies List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-list-all.json
     */
    /**
     * Sample code: Lists all Content Key Policies.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void listsAllContentKeyPolicies(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .contentKeyPolicies()
            .list("contoso", "contosomedia", null, null, null, com.azure.core.util.Context.NONE);
    }
}
