import com.azure.resourcemanager.compute.fluent.models.AvailabilitySetInner;

/** Samples for AvailabilitySets CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-07-01/examples/availabilitySetExamples/AvailabilitySet_Create.json
     */
    /**
     * Sample code: Create an availability set.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAnAvailabilitySet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getAvailabilitySets()
            .createOrUpdateWithResponse(
                "myResourceGroup",
                "myAvailabilitySet",
                new AvailabilitySetInner()
                    .withLocation("westus")
                    .withPlatformUpdateDomainCount(20)
                    .withPlatformFaultDomainCount(2),
                com.azure.core.util.Context.NONE);
    }
}
