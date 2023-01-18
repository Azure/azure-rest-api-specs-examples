/** Samples for ResourceGuards GetDefaultUpdateProtectedItemRequestsObject. */
public final class Main {
    /*
     * x-ms-original-file: specification/dataprotection/resource-manager/Microsoft.DataProtection/preview/2022-11-01-preview/examples/ResourceGuardCRUD/GetDefaultUpdateProtectedItemRequests.json
     */
    /**
     * Sample code: Get DefaultOperationsRequestObject.
     *
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getDefaultOperationsRequestObject(
        com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager
            .resourceGuards()
            .getDefaultUpdateProtectedItemRequestsObjectWithResponse(
                "SampleResourceGroup", "swaggerExample", "default", com.azure.core.util.Context.NONE);
    }
}
