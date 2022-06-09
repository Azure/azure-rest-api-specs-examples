```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.securityinsights.models.EntityQueriesKind;

/** Samples for EntityQueries List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2021-09-01-preview/examples/entityQueries/GetEntityQueries.json
     */
    /**
     * Sample code: Get all entity queries.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllEntityQueries(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.entityQueries().list("myRg", "myWorkspace", EntityQueriesKind.EXPANSION, Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-securityinsights_1.0.0-beta.1/sdk/securityinsights/azure-resourcemanager-securityinsights/README.md) on how to add the SDK to your project and authenticate.
