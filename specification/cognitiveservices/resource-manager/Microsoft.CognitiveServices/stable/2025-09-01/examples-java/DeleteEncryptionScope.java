
/**
 * Samples for EncryptionScopes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2025-09-01/examples/
     * DeleteEncryptionScope.json
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
