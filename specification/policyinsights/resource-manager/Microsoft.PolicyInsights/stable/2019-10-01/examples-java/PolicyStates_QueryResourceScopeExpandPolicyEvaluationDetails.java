import com.azure.core.util.Context;
import com.azure.resourcemanager.policyinsights.models.PolicyStatesResource;

/** Samples for PolicyStates ListQueryResultsForResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_QueryResourceScopeExpandPolicyEvaluationDetails.json
     */
    /**
     * Sample code: Query all policy states at resource scope and expand policyEvaluationDetails.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryAllPolicyStatesAtResourceScopeAndExpandPolicyEvaluationDetails(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyStates()
            .listQueryResultsForResource(
                PolicyStatesResource.LATEST,
                "subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.ClassicCompute/domainNames/myDomainName",
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                "PolicyEvaluationDetails",
                null,
                Context.NONE);
    }

    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_QueryResourceScopeExpandComponentsGroupByWithAggregate.json
     */
    /**
     * Sample code: Query component policy compliance state count grouped by state type at resource scope filtered by
     * given assignment.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void
        queryComponentPolicyComplianceStateCountGroupedByStateTypeAtResourceScopeFilteredByGivenAssignment(
            com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyStates()
            .listQueryResultsForResource(
                PolicyStatesResource.LATEST,
                "subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/Vaults/myKVName",
                null,
                null,
                null,
                null,
                null,
                "policyAssignmentId eq"
                    + " '/subscriptions/e78961ba-36fe-4739-9212-e3031b4c8db7/providers/microsoft.authorization/policyassignments/560050f83dbb4a24974323f8'",
                null,
                "components($filter=ComplianceState eq 'NonCompliant' or ComplianceState eq"
                    + " 'Compliant';$apply=groupby((complianceState),aggregate($count as count)))",
                null,
                Context.NONE);
    }

    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_QuerySubscriptionLevelResourceScope.json
     */
    /**
     * Sample code: Query all policy states at subscription level resource scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryAllPolicyStatesAtSubscriptionLevelResourceScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyStates()
            .listQueryResultsForResource(
                PolicyStatesResource.DEFAULT,
                "subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/providers/Microsoft.SomeNamespace/someResourceType/someResourceName",
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                Context.NONE);
    }

    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_QueryResourceScope.json
     */
    /**
     * Sample code: Query all policy states at resource scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryAllPolicyStatesAtResourceScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyStates()
            .listQueryResultsForResource(
                PolicyStatesResource.DEFAULT,
                "subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.ClassicCompute/domainNames/myDomainName",
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                Context.NONE);
    }

    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_QueryNestedResourceScope.json
     */
    /**
     * Sample code: Query all policy states at nested resource scope.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void queryAllPolicyStatesAtNestedResourceScope(
        com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyStates()
            .listQueryResultsForResource(
                PolicyStatesResource.DEFAULT,
                "subscriptions/fff10b27-fff3-fff5-fff8-fffbe01e86a5/resourceGroups/myResourceGroup/providers/Microsoft.ServiceFabric/clusters/myCluster/applications/myApplication",
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                null,
                Context.NONE);
    }
}
