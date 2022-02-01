Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.3/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for WorkbookTemplates GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-11-20/examples/WorkbookTemplateGet.json
     */
    /**
     * Sample code: WorkbookTemplateGet.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workbookTemplateGet(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager
            .workbookTemplates()
            .getByResourceGroupWithResponse("my-resource-group", "my-resource-name", Context.NONE);
    }
}
```
