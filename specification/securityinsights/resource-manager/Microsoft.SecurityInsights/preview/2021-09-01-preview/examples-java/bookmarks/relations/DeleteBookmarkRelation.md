```java
import com.azure.core.util.Context;

/** Samples for BookmarkRelations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/bookmarks/relations/DeleteBookmarkRelation.json
     */
    /**
     * Sample code: Delete the bookmark relation.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void deleteTheBookmarkRelation(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .bookmarkRelations()
            .deleteWithResponse(
                "myRg",
                "myWorkspace",
                "2216d0e1-91e3-4902-89fd-d2df8c535096",
                "4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014",
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.1/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.
