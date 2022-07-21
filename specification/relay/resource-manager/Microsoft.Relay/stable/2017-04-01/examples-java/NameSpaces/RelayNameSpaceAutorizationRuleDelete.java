import com.azure.core.util.Context;

/** Samples for Namespaces DeleteAuthorizationRule. */
public final class Main {
    /*
     * x-ms-original-file: specification/relay/resource-manager/Microsoft.Relay/stable/2017-04-01/examples/NameSpaces/RelayNameSpaceAutorizationRuleDelete.json
     */
    /**
     * Sample code: RelayNameSpaceAutorizationRuleDelete.
     *
     * @param manager Entry point to RelayManager.
     */
    public static void relayNameSpaceAutorizationRuleDelete(com.azure.resourcemanager.relay.RelayManager manager) {
        manager
            .namespaces()
            .deleteAuthorizationRuleWithResponse(
                "resourcegroup", "example-RelayNamespace-01", "example-RelayAuthRules-01", Context.NONE);
    }
}
