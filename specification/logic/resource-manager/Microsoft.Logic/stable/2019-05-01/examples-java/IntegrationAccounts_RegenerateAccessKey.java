import com.azure.resourcemanager.logic.models.KeyType;
import com.azure.resourcemanager.logic.models.RegenerateActionParameter;

/** Samples for IntegrationAccounts RegenerateAccessKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccounts_RegenerateAccessKey.json
     */
    /**
     * Sample code: Regenerate access key.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void regenerateAccessKey(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccounts()
            .regenerateAccessKeyWithResponse(
                "testResourceGroup",
                "testIntegrationAccount",
                new RegenerateActionParameter().withKeyType(KeyType.PRIMARY),
                com.azure.core.util.Context.NONE);
    }
}
