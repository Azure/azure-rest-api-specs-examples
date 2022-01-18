Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.11.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.appservice.models.StaticSiteBuildProperties;
import com.azure.resourcemanager.appservice.models.StaticSitesWorkflowPreviewRequest;

/** Samples for StaticSites PreviewWorkflow. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/GenerateStaticSiteWorkflowPreview.json
     */
    /**
     * Sample code: Generates a preview workflow file for the static site.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void generatesAPreviewWorkflowFileForTheStaticSite(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getStaticSites()
            .previewWorkflowWithResponse(
                "West US 2",
                new StaticSitesWorkflowPreviewRequest()
                    .withRepositoryUrl("https://github.com/username/RepoName")
                    .withBranch("master")
                    .withBuildProperties(
                        new StaticSiteBuildProperties()
                            .withAppLocation("app")
                            .withApiLocation("api")
                            .withAppArtifactLocation("build")),
                Context.NONE);
    }
}
```
