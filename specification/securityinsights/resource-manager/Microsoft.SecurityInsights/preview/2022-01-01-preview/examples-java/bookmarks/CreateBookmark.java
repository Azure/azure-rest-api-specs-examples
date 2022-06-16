import com.azure.resourcemanager.securityinsights.models.AttackTactic;
import com.azure.resourcemanager.securityinsights.models.BookmarkEntityMappings;
import com.azure.resourcemanager.securityinsights.models.EntityFieldMapping;
import com.azure.resourcemanager.securityinsights.models.UserInfo;
import java.time.OffsetDateTime;
import java.util.Arrays;
import java.util.UUID;

/** Samples for Bookmarks CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/bookmarks/CreateBookmark.json
     */
    /**
     * Sample code: Creates or updates a bookmark.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesABookmark(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .bookmarks()
            .define("73e01a99-5cd7-4139-a149-9f2736ff2ab5")
            .withExistingWorkspace("myRg", "myWorkspace")
            .withEtag("\"0300bf09-0000-0000-0000-5c37296e0000\"")
            .withCreated(OffsetDateTime.parse("2021-09-01T13:15:30Z"))
            .withCreatedBy(new UserInfo().withObjectId(UUID.fromString("2046feea-040d-4a46-9e2b-91c2941bfa70")))
            .withDisplayName("My bookmark")
            .withLabels(Arrays.asList("Tag1", "Tag2"))
            .withNotes("Found a suspicious activity")
            .withQuery("SecurityEvent | where TimeGenerated > ago(1d) and TimeGenerated < ago(2d)")
            .withQueryResult("Security Event query result")
            .withUpdated(OffsetDateTime.parse("2021-09-01T13:15:30Z"))
            .withUpdatedBy(new UserInfo().withObjectId(UUID.fromString("2046feea-040d-4a46-9e2b-91c2941bfa70")))
            .withEntityMappings(
                Arrays
                    .asList(
                        new BookmarkEntityMappings()
                            .withEntityType("Account")
                            .withFieldMappings(
                                Arrays
                                    .asList(
                                        new EntityFieldMapping()
                                            .withIdentifier("Fullname")
                                            .withValue("johndoe@microsoft.com")))))
            .withTactics(Arrays.asList(AttackTactic.EXECUTION))
            .withTechniques(Arrays.asList("T1609"))
            .create();
    }
}
