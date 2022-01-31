Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.3/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for WorkItemConfigurations GetItem. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/WorkItemConfigGet.json
     */
    /**
     * Sample code: WorkItemConfigurationsGetDefault.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workItemConfigurationsGetDefault(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .workItemConfigurations()
            .getItemWithResponse("my-resource-group", "my-component", "Visual Studio Team Services", Context.NONE);
    }
}
```
