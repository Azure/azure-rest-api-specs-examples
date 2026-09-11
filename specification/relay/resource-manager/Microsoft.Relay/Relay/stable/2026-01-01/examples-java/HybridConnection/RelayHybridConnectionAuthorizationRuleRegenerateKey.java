
import com.azure.resourcemanager.relay.models.KeyType;
import com.azure.resourcemanager.relay.models.RegenerateAccessKeyParameters;

/**
 * Samples for HybridConnections RegenerateKeys.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/HybridConnection/RelayHybridConnectionAuthorizationRuleRegenerateKey.json
     */
    /**
     * Sample code: RelayHybridConnectionAuthorizationRuleRegenerateKey.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void
        relayHybridConnectionAuthorizationRuleRegenerateKey(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.hybridConnections().regenerateKeysWithResponse("resourcegroup", "example-RelayNamespace-01",
            "example-Relay-Hybrid-01", "example-RelayAuthRules-01",
            new RegenerateAccessKeyParameters().withKeyType(KeyType.PRIMARY_KEY), com.azure.core.util.Context.NONE);
    }
}
