/** Samples for ContentKeyPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/Metadata/stable/2022-08-01/examples/content-key-policies-get-by-name.json
     */
    /**
     * Sample code: Get a Content Key Policy by name.
     *
     * @param manager Entry point to MediaServicesManager.
     */
    public static void getAContentKeyPolicyByName(
        com.azure.resourcemanager.mediaservices.MediaServicesManager manager) {
        manager
            .contentKeyPolicies()
            .getWithResponse("contoso", "contosomedia", "PolicyWithMultipleOptions", com.azure.core.util.Context.NONE);
    }
}
