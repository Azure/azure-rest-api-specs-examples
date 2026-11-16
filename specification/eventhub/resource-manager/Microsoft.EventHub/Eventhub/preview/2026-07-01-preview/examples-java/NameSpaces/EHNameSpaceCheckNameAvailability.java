
import com.azure.resourcemanager.eventhubs.models.CheckNameAvailabilityParameter;

/**
 * Samples for Namespaces CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/EHNameSpaceCheckNameAvailability.json
     */
    /**
     * Sample code: NamespacesCheckNameAvailability.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void namespacesCheckNameAvailability(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getNamespaces().checkNameAvailabilityWithResponse(
            new CheckNameAvailabilityParameter().withName("sdk-Namespace-8458"), com.azure.core.util.Context.NONE);
    }
}
