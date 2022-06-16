import com.azure.core.util.Context;

/** Samples for Secrets Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Secrets_Delete.json
     */
    /**
     * Sample code: Secrets_Delete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void secretsDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getSecrets().delete("RG", "profile1", "secret1", Context.NONE);
    }
}
