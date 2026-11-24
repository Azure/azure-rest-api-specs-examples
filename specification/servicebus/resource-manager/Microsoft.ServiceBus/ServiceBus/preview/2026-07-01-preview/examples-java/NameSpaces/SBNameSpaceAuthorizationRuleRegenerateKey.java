
import com.azure.resourcemanager.servicebus.models.KeyType;
import com.azure.resourcemanager.servicebus.models.RegenerateAccessKeyParameters;

/**
 * Samples for Namespaces RegenerateKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/SBNameSpaceAuthorizationRuleRegenerateKey.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleRegenerateKey.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void
        nameSpaceAuthorizationRuleRegenerateKey(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getNamespaces().regenerateKeysWithResponse("ArunMonocle", "sdk-namespace-6914",
            "sdk-AuthRules-1788", new RegenerateAccessKeyParameters().withKeyType(KeyType.PRIMARY_KEY),
            com.azure.core.util.Context.NONE);
    }
}
