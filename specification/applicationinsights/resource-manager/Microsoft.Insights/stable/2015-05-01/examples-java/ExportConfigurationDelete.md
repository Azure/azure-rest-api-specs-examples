Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.2/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for ExportConfigurations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ExportConfigurationDelete.json
     */
    /**
     * Sample code: ExportConfigurationDelete.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void exportConfigurationDelete(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .exportConfigurations()
            .deleteWithResponse("my-resource-group", "my-component", "uGOoki0jQsyEs3IdQ83Q4QsNr4=", Context.NONE);
    }
}
```
