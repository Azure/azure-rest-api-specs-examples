
/**
 * Samples for CodeSigningAccounts GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-02-05-preview/CodeSigningAccounts_Get.json
     */
    /**
     * Sample code: Get a Trusted Signing Account.
     * 
     * @param manager Entry point to TrustedSigningManager.
     */
    public static void
        getATrustedSigningAccount(com.azure.resourcemanager.trustedsigning.TrustedSigningManager manager) {
        manager.codeSigningAccounts().getByResourceGroupWithResponse("MyResourceGroup", "MyAccount",
            com.azure.core.util.Context.NONE);
    }
}
