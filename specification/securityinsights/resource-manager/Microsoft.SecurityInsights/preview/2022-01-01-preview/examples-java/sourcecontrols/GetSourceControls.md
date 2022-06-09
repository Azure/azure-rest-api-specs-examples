```java
import com.azure.core.util.Context;

/** Samples for SourceControlsOperation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/sourcecontrols/GetSourceControls.json
     */
    /**
     * Sample code: Get all source controls.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllSourceControls(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.sourceControlsOperations().list("myRg", "myWorkspace", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.3/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.
