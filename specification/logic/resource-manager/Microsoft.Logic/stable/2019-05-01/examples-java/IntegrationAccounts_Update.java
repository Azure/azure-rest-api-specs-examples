import com.azure.resourcemanager.logic.models.IntegrationAccount;
import com.azure.resourcemanager.logic.models.IntegrationAccountSku;
import com.azure.resourcemanager.logic.models.IntegrationAccountSkuName;

/** Samples for IntegrationAccounts Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccounts_Update.json
     */
    /**
     * Sample code: Patch an integration account.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void patchAnIntegrationAccount(com.azure.resourcemanager.logic.LogicManager manager) {
        IntegrationAccount resource =
            manager
                .integrationAccounts()
                .getByResourceGroupWithResponse(
                    "testResourceGroup", "testIntegrationAccount", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withSku(new IntegrationAccountSku().withName(IntegrationAccountSkuName.STANDARD)).apply();
    }
}
