import com.azure.core.util.Context;

/** Samples for ContentKeyPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mediaservices/resource-manager/Microsoft.Media/stable/2021-11-01/examples/content-key-policies-get-by-name.json
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
            .getWithResponse("contoso", "contosomedia", "PolicyWithMultipleOptions", Context.NONE);
    }
}
