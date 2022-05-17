Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.3/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for SourceControlsOperation Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/sourcecontrols/DeleteSourceControl.json
     */
    /**
     * Sample code: Delete a source control.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteASourceControl(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .sourceControlsOperations()
            .deleteWithResponse("myRg", "myWorkspace", "789e0c1f-4a3d-43ad-809c-e713b677b04a", Context.NONE);
    }
}
```
