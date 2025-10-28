
/**
 * Samples for AccountConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * AccountConnection/list.json
     */
    /**
     * Sample code: ListAccountConnections.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listAccountConnections(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accountConnections().list("resourceGroup-1", "account-1", "[tartget url]", "ContainerRegistry", null,
            com.azure.core.util.Context.NONE);
    }
}
