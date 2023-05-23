import com.azure.resourcemanager.frontdoor.models.CheckNameAvailabilityInput;
import com.azure.resourcemanager.frontdoor.models.ResourceType;

/** Samples for FrontDoorNameAvailabilityWithSubscription Check. */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/CheckFrontdoorNameAvailabilityWithSubscription.json
     */
    /**
     * Sample code: CheckNameAvailabilityWithSubscription.
     *
     * @param manager Entry point to FrontDoorManager.
     */
    public static void checkNameAvailabilityWithSubscription(
        com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager
            .frontDoorNameAvailabilityWithSubscriptions()
            .checkWithResponse(
                new CheckNameAvailabilityInput()
                    .withName("sampleName")
                    .withType(ResourceType.MICROSOFT_NETWORK_FRONT_DOORS_FRONTEND_ENDPOINTS),
                com.azure.core.util.Context.NONE);
    }
}
