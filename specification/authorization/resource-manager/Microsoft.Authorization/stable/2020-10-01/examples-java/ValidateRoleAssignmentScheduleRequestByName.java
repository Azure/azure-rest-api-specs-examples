import com.azure.core.util.Context;
import com.azure.resourcemanager.authorization.fluent.models.RoleAssignmentScheduleRequestInner;
import com.azure.resourcemanager.authorization.models.RequestType;
import com.azure.resourcemanager.authorization.models.RoleAssignmentScheduleRequestPropertiesScheduleInfo;
import com.azure.resourcemanager.authorization.models.RoleAssignmentScheduleRequestPropertiesScheduleInfoExpiration;
import com.azure.resourcemanager.authorization.models.Type;
import java.time.OffsetDateTime;

/** Samples for RoleAssignmentScheduleRequests Validate. */
public final class Main {
    /*
     * x-ms-original-file: specification/authorization/resource-manager/Microsoft.Authorization/stable/2020-10-01/examples/ValidateRoleAssignmentScheduleRequestByName.json
     */
    /**
     * Sample code: ValidateRoleAssignmentScheduleRequestByName.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void validateRoleAssignmentScheduleRequestByName(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .accessManagement()
            .roleAssignments()
            .manager()
            .roleServiceClient()
            .getRoleAssignmentScheduleRequests()
            .validateWithResponse(
                "subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f",
                "fea7a502-9a96-4806-a26f-eee560e52045",
                new RoleAssignmentScheduleRequestInner()
                    .withRoleDefinitionId(
                        "/subscriptions/dfa2a084-766f-4003-8ae1-c4aeb893a99f/providers/Microsoft.Authorization/roleDefinitions/c8d4ff99-41c3-41a8-9f60-21dfdad59608")
                    .withPrincipalId("a3bb8764-cb92-4276-9d2a-ca1e895e55ea")
                    .withRequestType(RequestType.SELF_ACTIVATE)
                    .withScheduleInfo(
                        new RoleAssignmentScheduleRequestPropertiesScheduleInfo()
                            .withStartDateTime(OffsetDateTime.parse("2020-09-09T21:35:27.91Z"))
                            .withExpiration(
                                new RoleAssignmentScheduleRequestPropertiesScheduleInfoExpiration()
                                    .withType(Type.AFTER_DURATION)
                                    .withDuration("PT8H")))
                    .withLinkedRoleEligibilityScheduleId("b1477448-2cc6-4ceb-93b4-54a202a89413")
                    .withCondition(
                        "@Resource[Microsoft.Storage/storageAccounts/blobServices/containers:ContainerName]"
                            + " StringEqualsIgnoreCase 'foo_storage_container'")
                    .withConditionVersion("1.0"),
                Context.NONE);
    }
}
