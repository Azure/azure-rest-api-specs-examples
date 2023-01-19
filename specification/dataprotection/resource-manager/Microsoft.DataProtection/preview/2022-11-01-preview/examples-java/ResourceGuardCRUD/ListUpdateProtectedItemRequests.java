/** Samples for ResourceGuards GetUpdateProtectedItemRequestsObjects. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/ResourceGuardCRUD/ListUpdateProtectedItemRequests.json
     */
    /**
     * Sample code: List OperationsRequestObject.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void listOperationsRequestObject(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager
            .resourceGuards()
            .getUpdateProtectedItemRequestsObjects(
                "SampleResourceGroup", "swaggerExample", com.azure.core.util.Context.NONE);
    }
}
