
import com.azure.resourcemanager.digitaltwins.models.CheckNameRequest;

/**
 * Samples for DigitalTwins CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/stable/2023-01-31/examples/
     * DigitalTwinsCheckNameAvailability_example.json
     */
    /**
     * Sample code: Check name Availability.
     * 
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void checkNameAvailability(com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager.digitalTwins().checkNameAvailabilityWithResponse("WestUS2",
            new CheckNameRequest().withName("myadtinstance"), com.azure.core.util.Context.NONE);
    }
}
