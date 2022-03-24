Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.2/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.securityinsights.models.ContentPathMap;
import com.azure.resourcemanager.securityinsights.models.ContentType;
import com.azure.resourcemanager.securityinsights.models.RepoType;
import com.azure.resourcemanager.securityinsights.models.Repository;
import java.util.Arrays;

/** Samples for SourceControlsOperation Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/sourcecontrols/CreateSourceControl.json
     */
    /**
     * Sample code: Creates a source control.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsASourceControl(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .sourceControlsOperations()
            .define("789e0c1f-4a3d-43ad-809c-e713b677b04a")
            .withExistingWorkspace("myRg", "myWorkspace")
            .withEtag("\"0300bf09-0000-0000-0000-5c37296e0000\"")
            .withDisplayName("My Source Control")
            .withDescription("This is a source control")
            .withRepoType(RepoType.GITHUB)
            .withContentTypes(Arrays.asList(ContentType.fromString("AnalyticRules"), ContentType.WORKBOOK))
            .withRepository(
                new Repository()
                    .withUrl("https://github.com/user/repo")
                    .withBranch("master")
                    .withDisplayUrl("https://github.com/user/repo")
                    .withPathMapping(
                        Arrays
                            .asList(
                                new ContentPathMap()
                                    .withContentType(ContentType.fromString("AnalyticRules"))
                                    .withPath("path/to/rules"),
                                new ContentPathMap()
                                    .withContentType(ContentType.WORKBOOK)
                                    .withPath("path/to/workbooks"))))
            .create();
    }
}
```
