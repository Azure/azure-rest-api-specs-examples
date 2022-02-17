Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-digitaltwins_1.0.0-beta.2/sdk/digitaltwins/azure-resourcemanager-digitaltwins/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.digitaltwins.models.CheckNameRequest;

/** Samples for DigitalTwins CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/digitaltwins/resource-manager/Microsoft.DigitalTwins/preview/2021-06-30-preview/examples/DigitalTwinsCheckNameAvailability_example.json
     */
    /**
     * Sample code: Check name Availability.
     *
     * @param manager Entry point to AzureDigitalTwinsManager.
     */
    public static void checkNameAvailability(com.azure.resourcemanager.digitaltwins.AzureDigitalTwinsManager manager) {
        manager
            .digitalTwins()
            .checkNameAvailabilityWithResponse(
                "WestUS2", new CheckNameRequest().withName("myadtinstance"), Context.NONE);
    }
}
```
