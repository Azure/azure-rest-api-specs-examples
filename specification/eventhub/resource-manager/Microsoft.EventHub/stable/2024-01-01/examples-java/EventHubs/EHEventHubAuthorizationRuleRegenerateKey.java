
import com.azure.resourcemanager.eventhubs.models.KeyType;
import com.azure.resourcemanager.eventhubs.models.RegenerateAccessKeyParameters;

/**
 * Samples for EventHubs RegenerateKeys.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/EventHubs/
     * EHEventHubAuthorizationRuleRegenerateKey.json
     */
    /**
     * Sample code: EventHubAuthorizationRuleRegenerateKey.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void eventHubAuthorizationRuleRegenerateKey(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getEventHubs().regenerateKeysWithResponse("ArunMonocle",
            "sdk-namespace-960", "sdk-EventHub-532", "sdk-Authrules-1534",
            new RegenerateAccessKeyParameters().withKeyType(KeyType.PRIMARY_KEY), com.azure.core.util.Context.NONE);
    }
}
