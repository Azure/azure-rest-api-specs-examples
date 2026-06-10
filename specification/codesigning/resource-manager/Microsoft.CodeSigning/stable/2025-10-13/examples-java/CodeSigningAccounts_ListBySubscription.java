
/**
 * Samples for CodeSigningAccounts List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-13/CodeSigningAccounts_ListBySubscription.json
     */
    /**
     * Sample code: Lists trusted signing accounts within a subscription.
     * 
     * @param manager Entry point to TrustedSigningManager.
     */
    public static void listsTrustedSigningAccountsWithinASubscription(
        com.azure.resourcemanager.trustedsigning.TrustedSigningManager manager) {
        manager.codeSigningAccounts().list(com.azure.core.util.Context.NONE);
    }
}
