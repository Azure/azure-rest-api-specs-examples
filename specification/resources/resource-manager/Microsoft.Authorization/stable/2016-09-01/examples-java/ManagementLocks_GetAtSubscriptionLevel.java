import com.azure.core.util.Context;

/** Samples for ManagementLocks GetAtSubscriptionLevel. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2016-09-01/examples/ManagementLocks_GetAtSubscriptionLevel.json
     */
    /**
     * Sample code: Get management lock at subscription level.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getManagementLockAtSubscriptionLevel(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .managementLockClient()
            .getManagementLocks()
            .getAtSubscriptionLevelWithResponse("testlock", Context.NONE);
    }
}
