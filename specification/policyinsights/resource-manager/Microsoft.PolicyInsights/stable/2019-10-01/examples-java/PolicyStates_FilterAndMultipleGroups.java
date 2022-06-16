import com.azure.core.util.Context;
import com.azure.resourcemanager.policyinsights.models.PolicyStatesResource;

/** Samples for PolicyStates ListQueryResultsForSubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyStates_FilterAndMultipleGroups.json
     */
    /**
     * Sample code: Filter and multiple groups.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void filterAndMultipleGroups(com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyStates()
            .listQueryResultsForSubscription(
                PolicyStatesResource.LATEST,
                "fffedd8f-ffff-fffd-fffd-fffed2f84852",
                10,
                "NumNonCompliantResources desc",
                null,
                null,
                null,
                "IsCompliant eq false",
                "groupby((PolicyAssignmentId, PolicySetDefinitionId, PolicyDefinitionId, PolicyDefinitionReferenceId,"
                    + " ResourceId))/groupby((PolicyAssignmentId, PolicySetDefinitionId, PolicyDefinitionId,"
                    + " PolicyDefinitionReferenceId), aggregate($count as NumNonCompliantResources))",
                null,
                Context.NONE);
    }
}
