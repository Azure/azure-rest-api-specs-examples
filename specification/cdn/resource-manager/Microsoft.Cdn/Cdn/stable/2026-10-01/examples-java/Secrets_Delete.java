
/**
 * Samples for Secrets Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Secrets_Delete.json
     */
    /**
     * Sample code: Secrets_Delete.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void secretsDelete(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getSecrets().delete("RG", "profile1", "secret1", com.azure.core.util.Context.NONE);
    }
}
