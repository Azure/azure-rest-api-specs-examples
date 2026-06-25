
/**
 * Samples for AvailabilitySets List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/availabilitySetExamples/AvailabilitySet_ListBySubscription.json
     */
    /**
     * Sample code: List availability sets in a subscription.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listAvailabilitySetsInASubscription(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getAvailabilitySets().list("virtualMachines\\$ref", com.azure.core.util.Context.NONE);
    }
}
