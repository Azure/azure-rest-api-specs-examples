```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.TeamProperties;

/** Samples for Incidents CreateTeam. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/incidents/CreateTeam.json
     */
    /**
     * Sample code: Creates incident teams group.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsIncidentTeamsGroup(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .incidents()
            .createTeamWithResponse(
                "ambawolvese5resourcegroup",
                "AmbaE5WestCentralUS",
                "69a30280-6a4c-4aa7-9af0-5d63f335d600",
                new TeamProperties().withTeamName("Team name").withTeamDescription("Team description"),
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.1/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.
