import com.azure.core.util.Context;

/** Samples for Namespaces ListKeys. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/NameSpaces/RelayNameSpaceAuthorizationRuleListKey.json
     */
    /**
     * Sample code: RelayNameSpaceAuthorizationRuleListKey.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayNameSpaceAuthorizationRuleListKey(com.azure.resourcemanager.relay.RelayManager manager) {
        manager
            .namespaces()
            .listKeysWithResponse(
                "resourcegroup", "example-RelayNamespace-01", "example-RelayAuthRules-01", Context.NONE);
    }
}
