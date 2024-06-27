
/**
 * Samples for Features List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/listProviderFeatures.json
     */
    /**
     * Sample code: List provider Features.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listProviderFeatures(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().featureClient().getFeatures().list("Resource Provider Namespace",
            com.azure.core.util.Context.NONE);
    }
}
