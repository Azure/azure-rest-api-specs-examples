
/**
 * Samples for EncryptionScopes List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-06-01/examples/
     * ListEncryptionScopes.json
     */
    /**
     * Sample code: ListEncryptionScopes.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listEncryptionScopes(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.encryptionScopes().list("resourceGroupName", "accountName", com.azure.core.util.Context.NONE);
    }
}
