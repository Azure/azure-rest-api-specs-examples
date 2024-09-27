
import com.azure.resourcemanager.trustedsigning.models.AccountSku;
import com.azure.resourcemanager.trustedsigning.models.CodeSigningAccountProperties;
import com.azure.resourcemanager.trustedsigning.models.SkuName;

/**
 * Samples for CodeSigningAccounts Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-02-05-preview/CodeSigningAccounts_Create.json
     */
    /**
     * Sample code: Create a trusted Signing Account.
     * 
     * @param manager Entry point to TrustedSigningManager.
     */
    public static void
        createATrustedSigningAccount(com.azure.resourcemanager.trustedsigning.TrustedSigningManager manager) {
        manager.codeSigningAccounts().define("MyAccount").withRegion("westus")
            .withExistingResourceGroup("MyResourceGroup")
            .withProperties(new CodeSigningAccountProperties().withSku(new AccountSku().withName(SkuName.BASIC)))
            .create();
    }
}
