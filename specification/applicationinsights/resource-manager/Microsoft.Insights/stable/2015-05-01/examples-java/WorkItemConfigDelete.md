```java
import com.azure.core.util.Context;

/** Samples for WorkItemConfigurations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WorkItemConfigDelete.json
     */
    /**
     * Sample code: WorkItemConfigurationDelete.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workItemConfigurationDelete(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .workItemConfigurations()
            .deleteWithResponse("my-resource-group", "my-component", "Visual Studio Team Services", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.4/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.
