/** Samples for Tag AssignToOperation. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateApiOperationTag.json
     */
    /**
     * Sample code: ApiManagementCreateApiOperationTag.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateApiOperationTag(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .tags()
            .assignToOperationWithResponse(
                "rg1",
                "apimService1",
                "5931a75ae4bbd512a88c680b",
                "5931a75ae4bbd512a88c680a",
                "tagId1",
                com.azure.core.util.Context.NONE);
    }
}
