
import com.azure.resourcemanager.frontdoor.models.CheckNameAvailabilityInput;
import com.azure.resourcemanager.frontdoor.models.ResourceType;

/**
 * Samples for FrontDoorNameAvailability Check.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/
     * CheckFrontdoorNameAvailability.json
     */
    /**
     * Sample code: CheckNameAvailability.
     * 
     * @param manager Entry point to FrontDoorManager.
     */
    public static void checkNameAvailability(com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager.frontDoorNameAvailabilities().checkWithResponse(new CheckNameAvailabilityInput().withName("sampleName")
            .withType(ResourceType.MICROSOFT_NETWORK_FRONT_DOORS), com.azure.core.util.Context.NONE);
    }
}
