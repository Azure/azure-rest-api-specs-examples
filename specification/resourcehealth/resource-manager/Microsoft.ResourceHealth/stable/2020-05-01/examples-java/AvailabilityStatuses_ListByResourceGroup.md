Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-resourcehealth_1.0.0-beta.2/sdk/resourcehealth/azure-resourcemanager-resourcehealth/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for AvailabilityStatuses ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcehealth/resource-manager/Microsoft.ResourceHealth/stable/2020-05-01/examples/AvailabilityStatuses_ListByResourceGroup.json
     */
    /**
     * Sample code: ListByResourceGroup.
     *
     * @param manager Entry point to ResourceHealthManager.
     */
    public static void listByResourceGroup(com.azure.resourcemanager.resourcehealth.ResourceHealthManager manager) {
        manager
            .availabilityStatuses()
            .listByResourceGroup("resourceGroupName", null, "recommendedactions", Context.NONE);
    }
}
```
