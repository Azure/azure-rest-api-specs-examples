
import com.azure.resourcemanager.loganalytics.models.Tag;
import java.util.Arrays;

/**
 * Samples for SavedSearches CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/
     * WorkspacesSavedSearchesCreateOrUpdate.json
     */
    /**
     * Sample code: SavedSearchCreateOrUpdate.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void savedSearchCreateOrUpdate(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.savedSearches().define("00000000-0000-0000-0000-00000000000").withExistingWorkspace("TestRG", "TestWS")
            .withCategory("Saved Search Test Category").withDisplayName("Create or Update Saved Search Test")
            .withQuery("Heartbeat | summarize Count() by Computer | take a")
            .withTags(Arrays.asList(new Tag().withName("Group").withValue("Computer")))
            .withFunctionAlias("heartbeat_func").withFunctionParameters("a:int=1").withVersion(2L).create();
    }
}
