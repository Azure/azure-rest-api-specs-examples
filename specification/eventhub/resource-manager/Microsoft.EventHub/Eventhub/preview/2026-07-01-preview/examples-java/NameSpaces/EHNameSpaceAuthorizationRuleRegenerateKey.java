
import com.azure.resourcemanager.eventhubs.models.KeyType;
import com.azure.resourcemanager.eventhubs.models.RegenerateAccessKeyParameters;

/**
 * Samples for Namespaces RegenerateKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/EHNameSpaceAuthorizationRuleRegenerateKey.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleRegenerateKey.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void
        nameSpaceAuthorizationRuleRegenerateKey(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getNamespaces().regenerateKeysWithResponse("ArunMonocle", "sdk-Namespace-8980",
            "sdk-Authrules-8929", new RegenerateAccessKeyParameters().withKeyType(KeyType.PRIMARY_KEY),
            com.azure.core.util.Context.NONE);
    }
}
