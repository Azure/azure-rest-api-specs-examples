
/**
 * Samples for CodeSigningAccounts ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-13/CodeSigningAccounts_ListByResourceGroup.json
     */
    /**
     * Sample code: Lists trusted signing accounts within a resource group.
     * 
     * @param manager Entry point to TrustedSigningManager.
     */
    public static void listsTrustedSigningAccountsWithinAResourceGroup(
        com.azure.resourcemanager.trustedsigning.TrustedSigningManager manager) {
        manager.codeSigningAccounts().listByResourceGroup("MyResourceGroup", com.azure.core.util.Context.NONE);
    }
}
