
import com.azure.core.util.Context;
import com.azure.resourcemanager.eventhubs.models.KeyType;
import com.azure.resourcemanager.eventhubs.models.RegenerateAccessKeyParameters;

/** Samples for Namespaces RegenerateKeys. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/
     * EHNameSpaceAuthorizationRuleRegenerateKey.json
     */
    /**
     * Sample code: NameSpaceAuthorizationRuleRegenerateKey.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpaceAuthorizationRuleRegenerateKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getNamespaces().regenerateKeysWithResponse("ArunMonocle",
            "sdk-Namespace-8980", "sdk-Authrules-8929",
            new RegenerateAccessKeyParameters().withKeyType(KeyType.PRIMARY_KEY), Context.NONE);
    }
}
