
import com.azure.resourcemanager.compute.fluent.models.AvailabilitySetInner;

/**
 * Samples for AvailabilitySets CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/availabilitySetExamples/AvailabilitySet_Create.json
     */
    /**
     * Sample code: Create an availability set.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void createAnAvailabilitySet(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getAvailabilitySets()
            .createOrUpdateWithResponse(
                "myResourceGroup", "myAvailabilitySet", new AvailabilitySetInner().withLocation("westus")
                    .withPlatformUpdateDomainCount(20).withPlatformFaultDomainCount(2),
                com.azure.core.util.Context.NONE);
    }
}
