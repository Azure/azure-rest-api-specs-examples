
/**
 * Samples for ResourceGuards GetDefaultUpdateProtectionPolicyRequestsObject.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/dataprotection/resource-manager/Microsoft.DataProtection/stable/2024-04-01/examples/
     * ResourceGuardCRUD/GetDefaultUpdateProtectionPolicyRequests.json
     */
    /**
     * Sample code: Get DefaultOperationsRequestObject.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        getDefaultOperationsRequestObject(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.resourceGuards().getDefaultUpdateProtectionPolicyRequestsObjectWithResponse("SampleResourceGroup",
            "swaggerExample", "default", com.azure.core.util.Context.NONE);
    }
}
