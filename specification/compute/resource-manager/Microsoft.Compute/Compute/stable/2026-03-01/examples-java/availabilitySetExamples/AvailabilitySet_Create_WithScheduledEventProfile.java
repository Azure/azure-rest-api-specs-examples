
import com.azure.resourcemanager.compute.fluent.models.AvailabilitySetInner;

/**
 * Samples for AvailabilitySets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/availabilitySetExamples/AvailabilitySet_Create_WithScheduledEventProfile.json
     */
    /**
     * Sample code: Create an availability set with Scheduled Event Policy.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        createAnAvailabilitySetWithScheduledEventPolicy(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getAvailabilitySets()
            .createOrUpdateWithResponse(
                "myResourceGroup", "myAvailabilitySet", new AvailabilitySetInner().withLocation("westus")
                    .withPlatformUpdateDomainCount(20).withPlatformFaultDomainCount(2),
                com.azure.core.util.Context.NONE);
    }
}
