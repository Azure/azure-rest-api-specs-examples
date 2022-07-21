import com.azure.core.util.Context;
import com.azure.resourcemanager.relay.models.KeyType;
import com.azure.resourcemanager.relay.models.RegenerateAccessKeyParameters;

/** Samples for HybridConnections RegenerateKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/HybridConnection/RelayHybridConnectionAuthorizationRuleRegenrateKey.json
     */
    /**
     * Sample code: RelayHybridConnectionAuthorizationRuleRegenrateKey.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayHybridConnectionAuthorizationRuleRegenrateKey(
        com.azure.resourcemanager.relay.RelayManager manager) {
        manager
            .hybridConnections()
            .regenerateKeysWithResponse(
                "resourcegroup",
                "example-RelayNamespace-01",
                "example-Relay-Hybrid-01",
                "example-RelayAuthRules-01",
                new RegenerateAccessKeyParameters().withKeyType(KeyType.PRIMARY_KEY),
                Context.NONE);
    }
}
