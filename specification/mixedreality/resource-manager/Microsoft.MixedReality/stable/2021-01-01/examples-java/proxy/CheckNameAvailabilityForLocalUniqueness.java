import com.azure.resourcemanager.mixedreality.models.CheckNameAvailabilityRequest;

/** Samples for ResourceProvider CheckNameAvailabilityLocal. */
public final class Main {
    /*
     * x-ms-original-file: specification/mixedreality/resource-manager/Microsoft.MixedReality/stable/2021-01-01/examples/proxy/CheckNameAvailabilityForLocalUniqueness.json
     */
    /**
     * Sample code: CheckLocalNameAvailability.
     *
     * @param manager Entry point to MixedRealityManager.
     */
    public static void checkLocalNameAvailability(com.azure.resourcemanager.mixedreality.MixedRealityManager manager) {
        manager
            .resourceProviders()
            .checkNameAvailabilityLocalWithResponse(
                "eastus2euap",
                new CheckNameAvailabilityRequest()
                    .withName("MyAccount")
                    .withType("Microsoft.MixedReality/spatialAnchorsAccounts"),
                com.azure.core.util.Context.NONE);
    }
}
