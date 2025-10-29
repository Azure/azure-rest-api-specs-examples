
/**
 * Samples for AccountConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * AccountConnection/delete.json
     */
    /**
     * Sample code: DeleteAccountConnection.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        deleteAccountConnection(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accountConnections().deleteWithResponse("resourceGroup-1", "account-1", "connection-1",
            com.azure.core.util.Context.NONE);
    }
}
