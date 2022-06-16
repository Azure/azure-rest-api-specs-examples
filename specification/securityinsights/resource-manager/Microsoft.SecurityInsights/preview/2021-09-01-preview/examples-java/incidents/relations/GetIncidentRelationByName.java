import com.azure.core.util.Context;

/** Samples for IncidentRelations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/incidents/relations/GetIncidentRelationByName.json
     */
    /**
     * Sample code: Get an incident relation.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAnIncidentRelation(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .incidentRelations()
            .getWithResponse(
                "myRg",
                "myWorkspace",
                "afbd324f-6c48-459c-8710-8d1e1cd03812",
                "4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014",
                Context.NONE);
    }
}
