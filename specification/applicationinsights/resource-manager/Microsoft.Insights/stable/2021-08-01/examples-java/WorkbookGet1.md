```java
import com.azure.core.util.Context;

/** Samples for Workbooks GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2021-08-01/examples/WorkbookGet1.json
     */
    /**
     * Sample code: WorkbookGet1.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workbookGet1(com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .workbooks()
            .getByResourceGroupWithResponse(
                "my-resource-group", "deadb33f-5e0d-4064-8ebb-1a4ed0313eb2", null, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.4/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.
