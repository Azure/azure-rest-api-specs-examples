```java
import com.azure.core.util.Context;
import java.time.Duration;

/** Samples for Reports ListByTime. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetReportsByTime.json
     */
    /**
     * Sample code: ApiManagementGetReportsByTime.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetReportsByTime(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .reports()
            .listByTime(
                "rg1",
                "apimService1",
                "timestamp ge datetime'2017-06-01T00:00:00' and timestamp le datetime'2017-06-04T00:00:00'",
                Duration.parse("PT15M"),
                null,
                null,
                null,
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-apimanagement_1.0.0-beta.3/sdk/apimanagement/azure-resourcemanager-apimanagement/README.md) on how to add the SDK to your project and authenticate.
