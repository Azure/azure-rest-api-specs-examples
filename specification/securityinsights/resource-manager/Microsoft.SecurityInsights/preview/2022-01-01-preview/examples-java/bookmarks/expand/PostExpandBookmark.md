Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.3/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.BookmarkExpandParameters;
import java.time.OffsetDateTime;
import java.util.UUID;

/** Samples for BookmarkOperation Expand. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/bookmarks/expand/PostExpandBookmark.json
     */
    /**
     * Sample code: Expand an bookmark.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void expandAnBookmark(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .bookmarkOperations()
            .expandWithResponse(
                "myRg",
                "myWorkspace",
                "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
                new BookmarkExpandParameters()
                    .withEndTime(OffsetDateTime.parse("2020-01-24T17:21:00.000Z"))
                    .withExpansionId(UUID.fromString("27f76e63-c41b-480f-bb18-12ad2e011d49"))
                    .withStartTime(OffsetDateTime.parse("2019-12-25T17:21:00.000Z")),
                Context.NONE);
    }
}
```
