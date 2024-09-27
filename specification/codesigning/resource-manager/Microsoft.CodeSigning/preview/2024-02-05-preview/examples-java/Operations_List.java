
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-02-05-preview/Operations_List.json
     */
    /**
     * Sample code: List trusted signing account operations.
     * 
     * @param manager Entry point to TrustedSigningManager.
     */
    public static void
        listTrustedSigningAccountOperations(com.azure.resourcemanager.trustedsigning.TrustedSigningManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
