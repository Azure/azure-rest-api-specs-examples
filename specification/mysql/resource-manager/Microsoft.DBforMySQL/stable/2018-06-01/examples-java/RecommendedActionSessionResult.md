```java
import com.azure.core.util.Context;

/** Samples for LocationBasedRecommendedActionSessionsResult List. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/RecommendedActionSessionResult.json
     */
    /**
     * Sample code: RecommendedActionSessionResult.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void recommendedActionSessionResult(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .locationBasedRecommendedActionSessionsResults()
            .list("WestUS", "aaaabbbb-cccc-dddd-0000-111122223333", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mysql_1.0.2/sdk/mysql/azure-resourcemanager-mysql/README.md) on how to add the SDK to your project and authenticate.
