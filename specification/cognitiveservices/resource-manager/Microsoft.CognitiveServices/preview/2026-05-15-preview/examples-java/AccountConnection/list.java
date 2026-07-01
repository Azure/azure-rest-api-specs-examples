
/**
 * Samples for AccountConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/AccountConnection/list.json
     */
    /**
     * Sample code: ListAccountConnections.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listAccountConnections(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accountConnections().list("resourceGroup-1", "account-1", "[target url]", "ContainerRegistry", null,
            com.azure.core.util.Context.NONE);
    }
}
