```java
import com.azure.core.util.Context;

/** Samples for CustomLocations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/stable/2021-08-15/examples/CustomLocationsListBySubscription.json
     */
    /**
     * Sample code: List Custom Locations by subscription.
     *
     * @param manager Entry point to CustomLocationsManager.
     */
    public static void listCustomLocationsBySubscription(
        com.azure.resourcemanager.extendedlocation.CustomLocationsManager manager) {
        manager.customLocations().list(Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-extendedlocation_1.0.0-beta.1/sdk/extendedlocation/azure-resourcemanager-extendedlocation/README.md) on how to add the SDK to your project and authenticate.
