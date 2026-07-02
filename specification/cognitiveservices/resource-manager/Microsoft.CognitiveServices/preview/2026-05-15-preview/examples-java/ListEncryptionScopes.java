
/**
 * Samples for EncryptionScopes List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-15-preview/ListEncryptionScopes.json
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
