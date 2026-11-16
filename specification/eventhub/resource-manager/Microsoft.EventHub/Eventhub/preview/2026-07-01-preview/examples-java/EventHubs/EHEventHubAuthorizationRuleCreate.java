
import com.azure.resourcemanager.eventhubs.fluent.models.AuthorizationRuleInner;
import com.azure.resourcemanager.eventhubs.models.AccessRights;
import java.util.Arrays;

/**
 * Samples for EventHubs CreateOrUpdateAuthorizationRule.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/EventHubs/EHEventHubAuthorizationRuleCreate.json
     */
    /**
     * Sample code: EventHubAuthorizationRuleCreate.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void eventHubAuthorizationRuleCreate(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getEventHubs().createOrUpdateAuthorizationRuleWithResponse("ArunMonocle",
            "sdk-Namespace-960", "sdk-EventHub-532", "sdk-Authrules-2513",
            new AuthorizationRuleInner().withRights(Arrays.asList(AccessRights.LISTEN, AccessRights.SEND)),
            com.azure.core.util.Context.NONE);
    }
}
