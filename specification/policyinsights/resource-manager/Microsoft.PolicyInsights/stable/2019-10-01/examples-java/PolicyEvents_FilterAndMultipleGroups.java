import com.azure.core.util.Context;
import com.azure.resourcemanager.policyinsights.models.PolicyEventsResourceType;
import java.time.OffsetDateTime;

/** Samples for PolicyEvents ListQueryResultsForSubscription. */
public final class Main {
    /*
     * x-ms-original-file: specification/policyinsights/resource-manager/Microsoft.PolicyInsights/stable/2019-10-01/examples/PolicyEvents_FilterAndMultipleGroups.json
     */
    /**
     * Sample code: Filter and multiple groups.
     *
     * @param manager Entry point to PolicyInsightsManager.
     */
    public static void filterAndMultipleGroups(com.azure.resourcemanager.policyinsights.PolicyInsightsManager manager) {
        manager
            .policyEvents()
            .listQueryResultsForSubscription(
                PolicyEventsResourceType.DEFAULT,
                "fffedd8f-ffff-fffd-fffd-fffed2f84852",
                10,
                "NumDeniedResources desc",
                null,
                OffsetDateTime.parse("2018-01-01T00:00:00Z"),
                null,
                "PolicyDefinitionAction eq 'deny'",
                "groupby((PolicyAssignmentId, PolicyDefinitionId, ResourceId))/groupby((PolicyAssignmentId,"
                    + " PolicyDefinitionId), aggregate($count as NumDeniedResources))",
                null,
                Context.NONE);
    }
}
