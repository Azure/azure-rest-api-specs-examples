import com.azure.resourcemanager.apimanagement.models.DocumentationContract;

/** Samples for Documentation Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateDocumentation.json
     */
    /**
     * Sample code: ApiManagementUpdateDocumentation.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateDocumentation(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        DocumentationContract resource =
            manager
                .documentations()
                .getWithResponse("rg1", "apimService1", "57d1f7558aa04f15146d9d8a", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().withTitle("Title updated").withContent("content updated").withIfMatch("*").apply();
    }
}
