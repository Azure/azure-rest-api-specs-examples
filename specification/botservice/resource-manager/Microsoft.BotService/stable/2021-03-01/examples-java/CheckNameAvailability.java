
import com.azure.resourcemanager.botservice.models.CheckNameAvailabilityRequestBody;

/**
 * Samples for Bots GetCheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/botservice/resource-manager/Microsoft.BotService/stable/2021-03-01/examples/CheckNameAvailability.
     * json
     */
    /**
     * Sample code: check Name Availability.
     * 
     * @param manager Entry point to BotServiceManager.
     */
    public static void checkNameAvailability(com.azure.resourcemanager.botservice.BotServiceManager manager) {
        manager.bots().getCheckNameAvailabilityWithResponse(
            new CheckNameAvailabilityRequestBody().withName("testbotname").withType("string"),
            com.azure.core.util.Context.NONE);
    }
}
