import com.azure.resourcemanager.apimanagement.models.WikiDocumentationContract;
import com.azure.resourcemanager.apimanagement.models.WikiUpdateContract;
import java.util.Arrays;

/** Samples for ApiWiki Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateApiWiki.json
     */
    /**
     * Sample code: ApiManagementUpdateApiWiki.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateApiWiki(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .apiWikis()
            .updateWithResponse(
                "rg1",
                "apimService1",
                "57d1f7558aa04f15146d9d8a",
                "*",
                new WikiUpdateContract()
                    .withDocuments(Arrays.asList(new WikiDocumentationContract().withDocumentationId("docId1"))),
                com.azure.core.util.Context.NONE);
    }
}
