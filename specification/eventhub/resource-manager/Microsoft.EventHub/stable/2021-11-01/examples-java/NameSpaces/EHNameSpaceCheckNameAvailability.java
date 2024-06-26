
import com.azure.core.util.Context;
import com.azure.resourcemanager.eventhubs.models.CheckNameAvailabilityParameter;

/** Samples for Namespaces CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/NameSpaces/
     * EHNameSpaceCheckNameAvailability.json
     */
    /**
     * Sample code: NamespacesCheckNameAvailability.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void namespacesCheckNameAvailability(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getNamespaces().checkNameAvailabilityWithResponse(
            new CheckNameAvailabilityParameter().withName("sdk-Namespace-8458"), Context.NONE);
    }
}
