
/**
 * Samples for AccountConnection Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/preview/2025-04-01-preview/examples/
     * AccountConnection/get.json
     */
    /**
     * Sample code: GetAccountConnection.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getAccountConnection(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accountConnections().getWithResponse("resourceGroup-1", "account-1", "connection-1",
            com.azure.core.util.Context.NONE);
    }
}
