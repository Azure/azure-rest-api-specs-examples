
/**
 * Samples for ResourceGuards GetDefaultUpdateProtectedItemRequestsObject.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/ResourceGuardCRUD/GetDefaultUpdateProtectedItemRequests.json
     */
    /**
     * Sample code: Get DefaultOperationsRequestObject.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        getDefaultOperationsRequestObject(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.resourceGuards().getDefaultUpdateProtectedItemRequestsObjectWithResponse("SampleResourceGroup",
            "swaggerExample", "default", com.azure.core.util.Context.NONE);
    }
}
