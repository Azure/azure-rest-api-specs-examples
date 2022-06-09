```java
import com.azure.core.util.Context;

/** Samples for ExportConfigurations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ExportConfigurationsList.json
     */
    /**
     * Sample code: ExportConfigurationsList.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void exportConfigurationsList(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.exportConfigurations().listWithResponse("my-resource-group", "my-component", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.4/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.
