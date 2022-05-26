Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-deviceupdate_1.0.0-beta.1/sdk/deviceupdate/azure-resourcemanager-deviceupdate/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.deviceupdate.models.CheckNameAvailabilityRequest;

/** Samples for ResourceProvider CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2022-04-01-preview/examples/CheckNameAvailability_Available.json
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
```
