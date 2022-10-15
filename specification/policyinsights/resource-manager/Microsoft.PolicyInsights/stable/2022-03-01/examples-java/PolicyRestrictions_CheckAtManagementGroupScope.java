import com.azure.core.util.Context;
import com.azure.resourcemanager.policyinsights.models.CheckManagementGroupRestrictionsRequest;
import com.azure.resourcemanager.policyinsights.models.PendingField;
import java.util.Arrays;

/** Samples for PolicyRestrictions CheckAtManagementGroupScope. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2022-03-01/examples/PolicyRestrictions_CheckAtManagementGroupScope.json
     */
    /**
     * Sample code: Check policy restrictions at management group scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void checkPolicyRestrictionsAtManagementGroupScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyRestrictions()
            .checkAtManagementGroupScopeWithResponse(
                "financeMg",
                new CheckManagementGroupRestrictionsRequest()
                    .withPendingFields(Arrays.asList(new PendingField().withField("type"))),
                Context.NONE);
    }
}
