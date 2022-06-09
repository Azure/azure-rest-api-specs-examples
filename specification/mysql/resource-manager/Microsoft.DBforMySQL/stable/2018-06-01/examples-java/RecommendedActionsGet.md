```java
import com.azure.core.util.Context;

/** Samples for RecommendedActions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/mysql/resource-manager/Microsoft.DBforMySQL/stable/2018-06-01/examples/RecommendedActionsGet.json
     */
    /**
     * Sample code: RecommendedActionsGet.
     *
     * @param manager Entry point to MySqlManager.
     */
    public static void recommendedActionsGet(com.azure.resourcemanager.mysql.MySqlManager manager) {
        manager
            .recommendedActions()
            .getWithResponse("testResourceGroupName", "testServerName", "Index", "Index-1", Context.NONE);
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-mysql_1.0.2/sdk/mysql/azure-resourcemanager-mysql/README.md) on how to add the SDK to your project and authenticate.
