import com.azure.core.util.Context;
import com.azure.resourcemanager.deviceupdate.models.CheckNameAvailabilityRequest;

/** Samples for ResourceProvider CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2022-04-01-preview/examples/CheckNameAvailability_AlreadyExists.json
     */
    /**
     * Sample code: CheckNameAvailability_AlreadyExists.
     *
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void checkNameAvailabilityAlreadyExists(
        com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager
            .resourceProviders()
            .checkNameAvailabilityWithResponse(
                new CheckNameAvailabilityRequest().withName("contoso").withType("Microsoft.DeviceUpdate/accounts"),
                Context.NONE);
    }
}
