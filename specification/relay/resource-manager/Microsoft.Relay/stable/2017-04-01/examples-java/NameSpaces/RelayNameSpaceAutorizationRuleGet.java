import com.azure.core.util.Context;

/** Samples for Namespaces GetAuthorizationRule. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/NameSpaces/RelayNameSpaceAutorizationRuleGet.json
     */
    /**
     * Sample code: RelayNameSpaceAutorizationRuleGet.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayNameSpaceAutorizationRuleGet(com.azure.resourcemanager.relay.RelayManager manager) {
        manager
            .namespaces()
            .getAuthorizationRuleWithResponse(
                "resourcegroup", "example-RelayNamespace-01", "example-RelayAuthRules-01", Context.NONE);
    }
}
