
/**
 * Samples for ResourceGuards GetDeleteResourceGuardProxyRequestsObjects.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/ResourceGuardCRUD/ListDeleteResourceGuardProxyRequests.json
     */
    /**
     * Sample code: List OperationsRequestObject.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        listOperationsRequestObject(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.resourceGuards().getDeleteResourceGuardProxyRequestsObjects("SampleResourceGroup", "swaggerExample",
            com.azure.core.util.Context.NONE);
    }
}
