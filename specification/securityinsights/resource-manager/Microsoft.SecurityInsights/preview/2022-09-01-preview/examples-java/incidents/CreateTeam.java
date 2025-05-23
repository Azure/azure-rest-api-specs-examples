
import com.azure.resourcemanager.securityinsights.models.TeamProperties;

/**
 * Samples for Incidents CreateTeam.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/
     * incidents/CreateTeam.json
     */
    /**
     * Sample code: Creates incident teams group.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        createsIncidentTeamsGroup(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.incidents().createTeamWithResponse("ambawolvese5resourcegroup", "AmbaE5WestCentralUS",
            "69a30280-6a4c-4aa7-9af0-5d63f335d600",
            new TeamProperties().withTeamName("Team name").withTeamDescription("Team description"),
            com.azure.core.util.Context.NONE);
    }
}
