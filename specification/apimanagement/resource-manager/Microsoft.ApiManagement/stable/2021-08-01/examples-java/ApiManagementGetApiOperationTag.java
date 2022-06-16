import com.azure.core.util.Context;

/** Samples for Tag GetByOperation. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetApiOperationTag.json
     */
    /**
     * Sample code: ApiManagementGetApiOperationTag.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetApiOperationTag(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .tags()
            .getByOperationWithResponse(
                "rg1",
                "apimService1",
                "59d6bb8f1f7fab13dc67ec9b",
                "59d6bb8f1f7fab13dc67ec9a",
                "59306a29e4bbd510dc24e5f9",
                Context.NONE);
    }
}
