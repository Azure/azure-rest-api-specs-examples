```java
import com.azure.core.util.Context;

/** Samples for EntityRelations GetRelation. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/entities/relations/GetEntityRelationByName.json
     */
    /**
     * Sample code: Get an entity relation.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAnEntityRelation(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager
            .entityRelations()
            .getRelationWithResponse(
                "myRg",
                "myWorkspace",
                "afbd324f-6c48-459c-8710-8d1e1cd03812",
                "4bb36b7b-26ff-4d1c-9cbe-0d8ab3da0014",
                Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.3/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.
