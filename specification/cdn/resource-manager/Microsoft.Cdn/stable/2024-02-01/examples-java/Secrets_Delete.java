
/**
 * Samples for Secrets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Secrets_Delete.json
     */
    /**
     * Sample code: Secrets_Delete.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void secretsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getSecrets().delete("RG", "profile1", "secret1",
            com.azure.core.util.Context.NONE);
    }
}
