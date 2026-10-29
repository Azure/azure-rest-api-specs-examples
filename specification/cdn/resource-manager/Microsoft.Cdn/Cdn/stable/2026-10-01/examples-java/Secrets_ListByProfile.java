
/**
 * Samples for Secrets ListByProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Secrets_ListByProfile.json
     */
    /**
     * Sample code: Secrets_ListByProfile.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void secretsListByProfile(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getSecrets().listByProfile("RG", "profile1", com.azure.core.util.Context.NONE);
    }
}
