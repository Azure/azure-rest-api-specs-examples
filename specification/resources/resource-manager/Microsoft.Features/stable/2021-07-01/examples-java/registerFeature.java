
/** Samples for Features Register. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/registerFeature.json
     */
    /**
     * Sample code: Register feature.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registerFeature(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().featureClient().getFeatures()
            .registerWithResponse("Resource Provider Namespace", "feature", com.azure.core.util.Context.NONE);
    }
}
