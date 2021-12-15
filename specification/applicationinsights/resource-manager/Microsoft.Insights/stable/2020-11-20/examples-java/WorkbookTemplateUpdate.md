Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.2/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.applicationinsights.models.WorkbookTemplate;

/** Samples for WorkbookTemplates Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2020-11-20/examples/WorkbookTemplateUpdate.json
     */
    /**
     * Sample code: WorkbookTemplateUpdate.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void workbookTemplateUpdate(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        WorkbookTemplate resource =
            manager
                .workbookTemplates()
                .getByResourceGroupWithResponse("my-resource-group", "my-template-resource", Context.NONE)
                .getValue();
        resource.update().apply();
    }
}
```
