Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.3/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;

/** Samples for Incidents ListEntities. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/incidents/entities/GetAllIncidentEntities.json
     */
    /**
     * Sample code: Gets all incident related entities.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getsAllIncidentRelatedEntities(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .incidents()
            .listEntitiesWithResponse("myRg", "myWorkspace", "afbd324f-6c48-459c-8710-8d1e1cd03812", Context.NONE);
    }
}
```
