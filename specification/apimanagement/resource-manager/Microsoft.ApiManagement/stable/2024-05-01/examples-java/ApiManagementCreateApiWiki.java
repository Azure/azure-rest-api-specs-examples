
import com.azure.resourcemanager.apimanagement.fluent.models.WikiContractInner;
import com.azure.resourcemanager.apimanagement.models.WikiDocumentationContract;
import java.util.Arrays;

/**
 * Samples for ApiWiki CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateApiWiki.json
     */
    /**
     * Sample code: ApiManagementCreateApiWiki.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateApiWiki(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiWikis().createOrUpdateWithResponse("rg1", "apimService1", "57d1f7558aa04f15146d9d8a",
            new WikiContractInner()
                .withDocuments(Arrays.asList(new WikiDocumentationContract().withDocumentationId("docId1"),
                    new WikiDocumentationContract().withDocumentationId("docId2"))),
            null, com.azure.core.util.Context.NONE);
    }
}
