
import com.azure.resourcemanager.eventhubs.models.KeyType;
import com.azure.resourcemanager.eventhubs.models.RegenerateAccessKeyParameters;

/**
 * Samples for EventHubs RegenerateKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/EventHubs/EHEventHubAuthorizationRuleRegenerateKey.json
     */
    /**
     * Sample code: EventHubAuthorizationRuleRegenerateKey.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void
        eventHubAuthorizationRuleRegenerateKey(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getEventHubs().regenerateKeysWithResponse("ArunMonocle", "sdk-namespace-960",
            "sdk-EventHub-532", "sdk-Authrules-1534",
            new RegenerateAccessKeyParameters().withKeyType(KeyType.PRIMARY_KEY), com.azure.core.util.Context.NONE);
    }
}
