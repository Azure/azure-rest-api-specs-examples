Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.1/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.securityinsights.models.Kind;

/** Samples for Metadata Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/metadata/PutMetadataMinimal.json
     */
    /**
     * Sample code: Create/update minimal metadata.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createUpdateMinimalMetadata(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .metadatas()
            .define("metadataName")
            .withExistingWorkspace("myRg", "myWorkspace")
            .withContentId("c00ee137-7475-47c8-9cce-ec6f0f1bedd0")
            .withParentId(
                "/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/ruleName")
            .withKind(Kind.ANALYTICS_RULE)
            .create();
    }
}
```
