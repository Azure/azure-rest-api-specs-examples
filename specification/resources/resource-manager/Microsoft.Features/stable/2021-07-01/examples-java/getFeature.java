
/** Samples for Features Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/getFeature.json
     */
    /**
     * Sample code: Get feature.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getFeature(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().featureClient().getFeatures().getWithResponse("Resource Provider Namespace",
            "feature", com.azure.core.util.Context.NONE);
    }
}
