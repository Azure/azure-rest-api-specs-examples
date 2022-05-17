Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.3/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for BookmarkRelations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/bookmarks/relations/GetAllBookmarkRelations.json
     */
    /**
     * Sample code: Get all bookmark relations.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllBookmarkRelations(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .bookmarkRelations()
            .list("myRg", "myWorkspace", "2216d0e1-91e3-4902-89fd-d2df8c535096", null, null, null, null, Context.NONE);
    }
}
```
