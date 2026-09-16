
import com.azure.resourcemanager.securityinsights.models.BookmarkExpandParameters;
import java.time.OffsetDateTime;

/**
 * Samples for BookmarkOperation Expand.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/bookmarks/expand/PostExpandBookmark.json
     */
    /**
     * Sample code: Expand an bookmark.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void expandAnBookmark(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.bookmarkOperations().expandWithResponse("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
            new BookmarkExpandParameters().withEndTime(OffsetDateTime.parse("2020-01-24T17:21:00.000Z"))
                .withExpansionId("27f76e63-c41b-480f-bb18-12ad2e011d49")
                .withStartTime(OffsetDateTime.parse("2019-12-25T17:21:00.000Z")),
            com.azure.core.util.Context.NONE);
    }
}
