
/** Samples for AccessPolicyAssignment Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/redis/resource-manager/Microsoft.Cache/stable/2023-08-01/examples/
     * RedisCacheAccessPolicyAssignmentGet.json
     */
    /**
     * Sample code: RedisCacheAccessPolicyAssignmentGet.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void redisCacheAccessPolicyAssignmentGet(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.redisCaches().manager().serviceClient().getAccessPolicyAssignments().getWithResponse("rg1", "cache1",
            "accessPolicyAssignmentName1", com.azure.core.util.Context.NONE);
    }
}
