
/**
 * Samples for PrivateLinkResources List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/ListPrivateLinkResources.json
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
