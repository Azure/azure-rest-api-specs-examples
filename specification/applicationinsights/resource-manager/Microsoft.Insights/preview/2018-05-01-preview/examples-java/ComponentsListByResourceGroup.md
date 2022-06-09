```java
import com.azure.core.util.Context;

/** Samples for Components ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2018-05-01-preview/examples/ComponentsListByResourceGroup.json
     */
    /**
     * Sample code: ComponentListByResourceGroup.
     *
     * @param manager Entry point to ApplicationInsightsManager.
     */
    public static void componentListByResourceGroup(
        com.azure.resourcemanager.applicationinsights.ApplicationInsightsManager manager) {
        manager.components().listByResourceGroup("my-resource-group", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-applicationinsights_1.0.0-beta.2/sdk/applicationinsights/azure-resourcemanager-applicationinsights/README.md) on how to add the SDK to your project and authenticate.
