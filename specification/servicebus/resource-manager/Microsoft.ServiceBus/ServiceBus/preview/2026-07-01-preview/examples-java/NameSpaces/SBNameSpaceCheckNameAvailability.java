
import com.azure.resourcemanager.servicebus.models.CheckNameAvailability;

/**
 * Samples for Namespaces CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/NameSpaces/SBNameSpaceCheckNameAvailability.json
     */
    /**
     * Sample code: NameSpaceCheckNameAvailability.
     * 
     * @param manager Entry point to ServiceBusManager.
     */
    public static void nameSpaceCheckNameAvailability(com.azure.resourcemanager.servicebus.ServiceBusManager manager) {
        manager.serviceClient().getNamespaces().checkNameAvailabilityWithResponse(
            new CheckNameAvailability().withName("sdk-Namespace-2924"), com.azure.core.util.Context.NONE);
    }
}
