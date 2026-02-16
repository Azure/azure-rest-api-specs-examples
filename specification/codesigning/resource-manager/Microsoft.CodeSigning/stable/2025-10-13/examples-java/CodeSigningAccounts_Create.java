
import com.azure.resourcemanager.artifactsigning.models.AccountSku;
import com.azure.resourcemanager.artifactsigning.models.SkuName;

/**
 * Samples for CodeSigningAccounts Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-13/CodeSigningAccounts_Create.json
     */
    /**
     * Sample code: Create an artifact signing account.
     * 
     * @param manager Entry point to ArtifactSigningManager.
     */
    public static void
        createAnArtifactSigningAccount(com.azure.resourcemanager.artifactsigning.ArtifactSigningManager manager) {
        manager.codeSigningAccounts().define("MyAccount").withRegion("westus")
            .withExistingResourceGroup("MyResourceGroup").withSku(new AccountSku().withName(SkuName.BASIC)).create();
    }
}
