
/**
 * Samples for Secrets Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Secrets_Get.json
     */
    /**
     * Sample code: Secrets_Get.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void secretsGet(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getSecrets().getWithResponse("RG", "profile1", "secret1",
            com.azure.core.util.Context.NONE);
    }
}
