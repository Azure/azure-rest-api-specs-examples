/** Samples for Tag Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteTag.json
     */
    /**
     * Sample code: ApiManagementDeleteTag.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteTag(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.tags().deleteWithResponse("rg1", "apimService1", "tagId1", "*", com.azure.core.util.Context.NONE);
    }
}
