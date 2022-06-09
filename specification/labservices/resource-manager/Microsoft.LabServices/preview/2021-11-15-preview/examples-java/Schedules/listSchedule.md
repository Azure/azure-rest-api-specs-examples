```java
import com.azure.core.util.Context;

/** Samples for Schedules ListByLab. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/Schedules/listSchedule.json
     */
    /**
     * Sample code: getListSchedule.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void getListSchedule(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.schedules().listByLab("testrg123", "testlab", null, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-labservices_1.0.0-beta.2/sdk/labservices/azure-resourcemanager-labservices/README.md) on how to add the SDK to your project and authenticate.
