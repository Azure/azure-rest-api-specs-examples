Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.3/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.resourcemanager.securityinsights.models.IncidentClassification;
import com.azure.resourcemanager.securityinsights.models.IncidentClassificationReason;
import com.azure.resourcemanager.securityinsights.models.IncidentOwnerInfo;
import com.azure.resourcemanager.securityinsights.models.IncidentSeverity;
import com.azure.resourcemanager.securityinsights.models.IncidentStatus;
import java.time.OffsetDateTime;
import java.util.UUID;

/** Samples for Incidents CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/incidents/CreateIncident.json
     */
    /**
     * Sample code: Creates or updates an incident.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesAnIncident(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .incidents()
            .define("73e01a99-5cd7-4139-a149-9f2736ff2ab5")
            .withExistingWorkspace("myRg", "myWorkspace")
            .withEtag("\"0300bf09-0000-0000-0000-5c37296e0000\"")
            .withClassification(IncidentClassification.FALSE_POSITIVE)
            .withClassificationComment("Not a malicious activity")
            .withClassificationReason(IncidentClassificationReason.INCORRECT_ALERT_LOGIC)
            .withDescription("This is a demo incident")
            .withFirstActivityTimeUtc(OffsetDateTime.parse("2019-01-01T13:00:30Z"))
            .withLastActivityTimeUtc(OffsetDateTime.parse("2019-01-01T13:05:30Z"))
            .withOwner(new IncidentOwnerInfo().withObjectId(UUID.fromString("2046feea-040d-4a46-9e2b-91c2941bfa70")))
            .withSeverity(IncidentSeverity.HIGH)
            .withStatus(IncidentStatus.CLOSED)
            .withTitle("My incident")
            .create();
    }
}
```
