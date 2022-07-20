import com.azure.core.util.Context;
import com.azure.resourcemanager.relay.models.KeyType;
import com.azure.resourcemanager.relay.models.RegenerateAccessKeyParameters;

/** Samples for WcfRelays RegenerateKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/Relay/RelayAuthorizationRuleRegenrateKey.json
     */
    /**
     * Sample code: RelayAuthorizationRuleRegenrateKey.json.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayAuthorizationRuleRegenrateKeyJson(com.azure.resourcemanager.relay.RelayManager manager) {
        manager
            .wcfRelays()
            .regenerateKeysWithResponse(
                "resourcegroup",
                "example-RelayNamespace-01",
                "example-Relay-wcf-01",
                "example-RelayAuthRules-01",
                new RegenerateAccessKeyParameters().withKeyType(KeyType.PRIMARY_KEY),
                Context.NONE);
    }
}
