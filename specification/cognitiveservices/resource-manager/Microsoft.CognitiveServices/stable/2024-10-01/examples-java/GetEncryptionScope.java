
/**
 * Samples for EncryptionScopes Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2024-10-01/examples/
     * GetEncryptionScope.json
     */
    /**
     * Sample code: GetEncryptionScope.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        getEncryptionScope(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.encryptionScopes().getWithResponse("resourceGroupName", "accountName", "encryptionScopeName",
            com.azure.core.util.Context.NONE);
    }
}
