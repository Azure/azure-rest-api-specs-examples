```java
import com.azure.core.util.Context;

/** Samples for Bookmarks Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/bookmarks/GetBookmarkById.json
     */
    /**
     * Sample code: Get a bookmark.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getABookmark(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .bookmarks()
            .getWithResponse("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.3/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.
