
/**
 * Samples for PrivateLinkResources List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/
     * ListPrivateLinkResources.json
     */
    /**
     * Sample code: ListPrivateLinkResources.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listPrivateLinkResources(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.privateLinkResources().listWithResponse("res6977", "sto2527", com.azure.core.util.Context.NONE);
    }
}
