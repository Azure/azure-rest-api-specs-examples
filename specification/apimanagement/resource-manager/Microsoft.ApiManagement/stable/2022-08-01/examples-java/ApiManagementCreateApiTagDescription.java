
/**
 * Samples for ApiTagDescription CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/
     * ApiManagementCreateApiTagDescription.json
     */
    /**
     * Sample code: ApiManagementCreateApiTagDescription.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateApiTagDescription(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiTagDescriptions().define("tagId1").withExistingApi("rg1", "apimService1", "5931a75ae4bbd512a88c680b")
            .withDescription(
                "Some description that will be displayed for operation's tag if the tag is assigned to operation of the API")
            .withExternalDocsUrl("http://some.url/additionaldoc")
            .withExternalDocsDescription("Description of the external docs resource").create();
    }
}
