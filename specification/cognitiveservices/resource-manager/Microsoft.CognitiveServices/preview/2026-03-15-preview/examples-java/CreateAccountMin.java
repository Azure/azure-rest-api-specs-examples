
import com.azure.resourcemanager.cognitiveservices.models.AccountProperties;
import com.azure.resourcemanager.cognitiveservices.models.Identity;
import com.azure.resourcemanager.cognitiveservices.models.ResourceIdentityType;
import com.azure.resourcemanager.cognitiveservices.models.Sku;

/**
 * Samples for Accounts Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/CreateAccountMin.json
     */
    /**
     * Sample code: Create Account Min.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void createAccountMin(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().define("testCreate1").withExistingResourceGroup("myResourceGroup").withRegion("West US")
            .withProperties(new AccountProperties()).withKind("CognitiveServices").withSku(new Sku().withName("S0"))
            .withIdentity(new Identity().withType(ResourceIdentityType.SYSTEM_ASSIGNED)).create();
    }
}
