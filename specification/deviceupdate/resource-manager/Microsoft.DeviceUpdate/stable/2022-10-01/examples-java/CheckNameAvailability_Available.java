import com.azure.core.util.Context;
import com.azure.resourcemanager.deviceupdate.models.CheckNameAvailabilityRequest;

/** Samples for ResourceProvider CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/CheckNameAvailability_Available.json
     */
    /**
     * Sample code: CheckNameAvailability_Available.
     *
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void checkNameAvailabilityAvailable(
        com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager
            .resourceProviders()
            .checkNameAvailabilityWithResponse(
                new CheckNameAvailabilityRequest().withName("contoso").withType("Microsoft.DeviceUpdate/accounts"),
                Context.NONE);
    }
}
