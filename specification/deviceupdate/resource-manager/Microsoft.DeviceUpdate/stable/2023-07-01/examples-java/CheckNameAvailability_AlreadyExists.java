
import com.azure.resourcemanager.deviceupdate.models.CheckNameAvailabilityRequest;

/**
 * Samples for ResourceProvider CheckNameAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/
     * CheckNameAvailability_AlreadyExists.json
     */
    /**
     * Sample code: CheckNameAvailability_AlreadyExists.
     * 
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void
        checkNameAvailabilityAlreadyExists(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.resourceProviders().checkNameAvailabilityWithResponse(
            new CheckNameAvailabilityRequest().withName("contoso").withType("Microsoft.DeviceUpdate/accounts"),
            com.azure.core.util.Context.NONE);
    }
}
