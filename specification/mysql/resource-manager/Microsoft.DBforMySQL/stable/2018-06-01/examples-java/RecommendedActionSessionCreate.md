```java
import com.azure.core.util.Context;

/** Samples for ResourceProvider CreateRecommendedActionSession. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/RecommendedActionSessionCreate.json
     */
    /**
     * Sample code: RecommendedActionSessionCreate.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void recommendedActionSessionCreate(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .resourceProviders()
            .createRecommendedActionSession(
                "testResourceGroupName", "testServerName", "Index", "someDatabaseName", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mysql_1.0.2/sdk/mysql/azure-resourcemanager-mysql/README.md) on how to add the SDK to your project and authenticate.
