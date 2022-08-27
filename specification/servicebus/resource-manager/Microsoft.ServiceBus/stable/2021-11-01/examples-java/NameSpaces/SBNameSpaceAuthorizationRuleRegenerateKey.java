import com.azure.core.util.Context;
import com.azure.resourcemanager.servicebus.models.KeyType;
import com.azure.resourcemanager.servicebus.models.RegenerateAccessKeyParameters;

/** Samples for Namespaces RegenerateKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/SBNameSpaceAuthorizationRuleRegenerateKey.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleRegenerateKey.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpaceAuthorizationRuleRegenerateKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .serviceBusNamespaces()
            .manager()
            .serviceClient()
            .getNamespaces()
            .regenerateKeysWithResponse(
                "ArunMonocle",
                "sdk-namespace-6914",
                "sdk-AuthRules-1788",
                new RegenerateAccessKeyParameters().withKeyType(KeyType.PRIMARY_KEY),
                Context.NONE);
    }
}
