import com.azure.core.util.Context;
import com.azure.resourcemanager.relay.fluent.models.AuthorizationRuleInner;
import com.azure.resourcemanager.relay.models.AccessRights;
import java.util.Arrays;

/** Samples for HybridConnections CreateOrUpdateAuthorizationRule. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/HybridConnection/RelayHybridConnectionAuthorizationRuleCreate.json
     */
    /**
     * Sample code: RelayHybridConnectionAuthorizationRuleCreate.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayHybridConnectionAuthorizationRuleCreate(
        com.azure.resourcemanager.relay.RelayManager manager) {
        manager
            .hybridConnections()
            .createOrUpdateAuthorizationRuleWithResponse(
                "resourcegroup",
                "example-RelayNamespace-01",
                "example-Relay-Hybrid-01",
                "example-RelayAuthRules-01",
                new AuthorizationRuleInner().withRights(Arrays.asList(AccessRights.LISTEN, AccessRights.SEND)),
                Context.NONE);
    }
}
