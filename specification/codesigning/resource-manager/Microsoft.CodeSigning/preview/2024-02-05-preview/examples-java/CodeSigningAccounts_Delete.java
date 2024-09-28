
/**
 * Samples for CodeSigningAccounts Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-02-05-preview/CodeSigningAccounts_Delete.json
     */
    /**
     * Sample code: Delete a trusted signing account.
     * 
     * @param manager Entry point to TrustedSigningManager.
     */
    public static void
        deleteATrustedSigningAccount(com.azure.resourcemanager.trustedsigning.TrustedSigningManager manager) {
        manager.codeSigningAccounts().delete("MyResourceGroup", "MyAccount", com.azure.core.util.Context.NONE);
    }
}
