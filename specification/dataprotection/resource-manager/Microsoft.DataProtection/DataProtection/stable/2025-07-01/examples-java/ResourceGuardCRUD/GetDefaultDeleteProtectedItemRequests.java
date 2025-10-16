
/**
 * Samples for ResourceGuards GetDefaultDeleteProtectedItemRequestsObject.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01/ResourceGuardCRUD/GetDefaultDeleteProtectedItemRequests.json
     */
    /**
     * Sample code: Get DefaultOperationsRequestObject.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        getDefaultOperationsRequestObject(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.resourceGuards().getDefaultDeleteProtectedItemRequestsObjectWithResponse("SampleResourceGroup",
            "swaggerExample", "default", com.azure.core.util.Context.NONE);
    }
}
