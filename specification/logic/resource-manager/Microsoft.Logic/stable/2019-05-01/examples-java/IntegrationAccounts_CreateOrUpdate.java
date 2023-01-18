import com.azure.resourcemanager.logic.models.IntegrationAccountSku;
import com.azure.resourcemanager.logic.models.IntegrationAccountSkuName;

/** Samples for IntegrationAccounts CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccounts_CreateOrUpdate.json
     */
    /**
     * Sample code: Create or update an integration account.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void createOrUpdateAnIntegrationAccount(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccounts()
            .define("testIntegrationAccount")
            .withRegion("westus")
            .withExistingResourceGroup("testResourceGroup")
            .withSku(new IntegrationAccountSku().withName(IntegrationAccountSkuName.STANDARD))
            .create();
    }
}
