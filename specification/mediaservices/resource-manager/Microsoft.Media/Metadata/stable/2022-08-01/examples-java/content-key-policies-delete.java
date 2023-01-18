/** Samples for ContentKeyPolicies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-delete.json
     */
    /**
     * Sample code: Delete a Key Policy.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void deleteAKeyPolicy(com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .contentKeyPolicies()
            .deleteWithResponse(
                "contoso",
                "contosomedia",
                "PolicyWithPlayReadyOptionAndOpenRestriction",
                com.azure.core.util.Context.NONE);
    }
}
