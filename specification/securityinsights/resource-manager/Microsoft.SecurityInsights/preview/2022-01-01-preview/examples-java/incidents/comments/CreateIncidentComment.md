Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.2/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.

```java
/** Samples for IncidentComments CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/incidents/comments/CreateIncidentComment.json
     */
    /**
     * Sample code: Creates or updates an incident comment.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void createsOrUpdatesAnIncidentComment(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .incidentComments()
            .define("4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014")
            .withExistingIncident("myRg", "myWorkspace", "73e01a99-5cd7-4139-a149-9f2736ff2ab5")
            .withMessage("Some message")
            .create();
    }
}
```
