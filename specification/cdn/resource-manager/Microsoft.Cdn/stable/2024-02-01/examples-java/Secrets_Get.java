
/**
 * Samples for Secrets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Secrets_Get.json
     */
    /**
     * Sample code: Secrets_Get.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void secretsGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getSecrets().getWithResponse("RG", "profile1", "secret1",
            com.azure.core.util.Context.NONE);
    }
}
