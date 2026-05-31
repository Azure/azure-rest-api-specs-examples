
/**
 * Samples for EncryptionScopes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/DeleteEncryptionScope.json
     */
    /**
     * Sample code: DeleteEncryptionScope.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        deleteEncryptionScope(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.encryptionScopes().delete("resourceGroupName", "accountName", "encryptionScopeName",
            com.azure.core.util.Context.NONE);
    }
}
