
/**
 * Samples for Features ListAll.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/listSubscriptionFeatures.
     * json
     */
    /**
     * Sample code: List subscription Features.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listSubscriptionFeatures(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().featureClient().getFeatures().listAll(com.azure.core.util.Context.NONE);
    }
}
