Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.15.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for RoleAssignmentScheduleRequests Cancel. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/CancelRoleAssignmentScheduleRequestByName.json
     */
    /**
     * Sample code: CancelRoleAssignmentScheduleRequestByName.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void cancelRoleAssignmentScheduleRequestByName(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getRoleAssignmentScheduleRequests()
            .cancelWithResponse(
                "providers/Microsoft.Subscription/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f",
                "fea7a502-9a96-4806-a26f-eee560e52045",
                Context.NONE);
    }
}
```
