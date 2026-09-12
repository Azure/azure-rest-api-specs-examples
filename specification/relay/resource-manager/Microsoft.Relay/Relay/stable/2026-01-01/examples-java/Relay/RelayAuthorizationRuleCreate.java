
import com.azure.resourcemanager.relay.fluent.models.AuthorizationRuleInner;
import com.azure.resourcemanager.relay.models.AccessRights;
import java.util.Arrays;

/**
 * Samples for WcfRelays CreateOrUpdateAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/Relay/RelayAuthorizationRuleCreate.json
     */
    /**
     * Sample code: RelayAuthorizationRuleCreate.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void relayAuthorizationRuleCreate(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.wcfRelays().createOrUpdateAuthorizationRuleWithResponse("resourcegroup", "example-RelayNamespace-01",
            "example-Relay-wcf-01", "example-RelayAuthRules-01",
            new AuthorizationRuleInner().withRights(Arrays.asList(AccessRights.LISTEN, AccessRights.SEND)),
            com.azure.core.util.Context.NONE);
    }
}
