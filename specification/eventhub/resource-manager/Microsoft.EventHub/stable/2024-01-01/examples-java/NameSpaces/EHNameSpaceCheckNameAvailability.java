
import com.azure.resourcemanager.eventhubs.models.CheckNameAvailabilityParameter;

/**
 * Samples for Namespaces CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/NameSpaces/
     * EHNameSpaceCheckNameAvailability.json
     */
    /**
     * Sample code: NamespacesCheckNameAvailability.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void namespacesCheckNameAvailability(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getNamespaces().checkNameAvailabilityWithResponse(
            new CheckNameAvailabilityParameter().withName("sdk-Namespace-8458"), com.azure.core.util.Context.NONE);
    }
}
