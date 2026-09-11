
import com.azure.resourcemanager.relay.models.AccessRights;
import java.util.Arrays;

/**
 * Samples for Namespaces CreateOrUpdateAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/NameSpaces/RelayNameSpaceAuthorizationRuleCreate.json
     */
    /**
     * Sample code: RelayNameSpaceAuthorizationRuleCreate.
     * 
     * @param manager Entry point to RelayManager.
     */
    public static void relayNameSpaceAuthorizationRuleCreate(com.azure.resourcemanager.relay.RelayManager manager) {
        manager.namespaces().defineAuthorizationRule("example-RelayAuthRules-01")
            .withExistingNamespace("resourcegroup", "example-RelayNamespace-01")
            .withRights(Arrays.asList(AccessRights.LISTEN, AccessRights.SEND)).create();
    }
}
