
import com.azure.core.util.Context;

/** Samples for EventHubs DeleteAuthorizationRule. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/EventHubs/
     * EHEventHubAuthorizationRuleDelete.json
     */
    /**
     * Sample code: EventHubAuthorizationRuleDelete.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void eventHubAuthorizationRuleDelete(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getEventHubs().deleteAuthorizationRuleWithResponse("ArunMonocle",
            "sdk-Namespace-960", "sdk-EventHub-532", "sdk-Authrules-2513", Context.NONE);
    }
}
