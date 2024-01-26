
import com.azure.core.util.Context;
import com.azure.resourcemanager.servicebus.models.CheckNameAvailability;

/** Samples for Namespaces CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/servicebus/resource-manager/Microsoft.ServiceBus/stable/2021-11-01/examples/NameSpaces/
     * SBNameSpaceCheckNameAvailability.json
     */
    /**
     * Sample code: NameSpaceCheckNameAvailability.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void nameSpaceCheckNameAvailability(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.serviceBusNamespaces().manager().serviceClient().getNamespaces().checkNameAvailabilityWithResponse(
            new CheckNameAvailability().withName("sdk-Namespace-2924"), Context.NONE);
    }
}
