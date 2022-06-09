```java
import com.azure.core.util.Context;

/** Samples for CustomLocations ListEnabledResourceTypes. */
public final class Main {
    /*
     * x-ms-original-file: specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/stable/2021-08-15/examples/CustomLocationsListEnabledResourceTypes.json
     */
    /**
     * Sample code: Get Custom Location.
     *
     * @param manager Entry point to CustomLocationsManager.
     */
    public static void getCustomLocation(com.azure.resourcemanager.extendedlocation.CustomLocationsManager manager) {
        manager.customLocations().listEnabledResourceTypes("testresourcegroup", "customLocation01", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-extendedlocation_1.0.0-beta.1/sdk/extendedlocation/azure-resourcemanager-extendedlocation/README.md) on how to add the SDK to your project and authenticate.
