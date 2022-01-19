Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-labservices_1.0.0-beta.2/sdk/labservices/azure-resourcemanager-labservices/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Schedules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Schedules/deleteSchedule.json
     */
    /**
     * Sample code: deleteSchedule.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void deleteSchedule(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.schedules().delete("testrg123", "testlab", "schedule1", Context.NONE);
    }
}
```
