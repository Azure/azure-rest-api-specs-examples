import com.azure.core.util.Context;

/** Samples for Features Unregister. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/unregisterFeature.json
     */
    /**
     * Sample code: Register feature.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void registerFeature(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .featureClient()
            .getFeatures()
            .unregisterWithResponse("Resource Provider Namespace", "feature", Context.NONE);
    }
}
